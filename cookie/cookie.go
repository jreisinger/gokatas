// Package cookie writes and reads a (not signed and not encrypted) HTTP
// [cookie]. It is a small piece of data that a server sends to a user's web
// browser or other HTTP client. This way server can map HTTP traffic (which is
// stateless) to a specific client. Cookies are used for:
//
//   - Session management (e.g. logins, shopping carts)
//   - Personalization (e.g. user preferences, themes)
//   - Tracking (recording and analyzing user behavior)-:
//
// Based on https://www.alexedwards.net/blog/working-with-cookies-in-go
//
// Level: intermediate
// Topics: web server, net/http
//
// [cookie]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies
package cookie

import (
	"errors"
	"log"
	"net/http"
)

const Name = "ExampleCookie"

// Set adds a "Set-Cookie" header to the response and sends it to a client.
func Set(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     Name,
		Value:    "Example Cookie",
		MaxAge:   3600,
		Secure:   true, // only sent over HTTPS (except on localhost)
		HttpOnly: true, // inaccessible to JavaScript (prevents XSS)
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)

	w.Write([]byte("cookie set: " + cookie.String()))
}

// Show retrieves the cookie from the "Cookie" request header and sends it back
// to the client in the response body.
func Show(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(Name)
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, "no cookie", http.StatusBadRequest)
		default:
			log.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Write([]byte("cookie: " + cookie.String()))
}
