// Package cookie writes and reads (not signed and not encrypted) [cookie]. An
// HTTP cookie is a small piece of data that a server sends to a user's web
// browser. Cookies are used mainly for:
//
//   - Session management (e.g. logins, shopping carts)
//   - Personalization (e.g. user preferences, themes)
//   - Tracking (recording and analyzing user behavior)
//
// Based on https://www.alexedwards.net/blog/working-with-cookies-in-go
//
// Level: intermediate
// Topics: net/http
//
// [cookies]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies
package cookie

import (
	"errors"
	"log"
	"net/http"
)

var (
	ErrValueTooLong = errors.New("cookie value too long")
	ErrInvalidValue = errors.New("invalid cookie value")
	name            = "exampleCookie"
)

// Set sets a cookie and sends it to a client.
func Set(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     name,
		Value:    "Hello world!",
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	w.Write([]byte("Set a cookie!"))
}

// Get retrieves the cookie from the request and sends it back to the client in
// the response body.
func Get(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(name)
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, "cookie not found", http.StatusBadRequest)
		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	}
	w.Write([]byte("Got back the cookie named " + cookie.Name))
}
