// Package cookie writes and reads (not signed and not encrypted) cookie. An
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
package cookie

import (
	"errors"
	"log"
	"net/http"
)

const Name = "exampleCookie"

// Set sets a cookie and sends it to a client in the response header.
func Set(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     Name,
		Value:    "hello world",
		MaxAge:   3600,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	w.Write([]byte("cookie set"))
}

// Show retrieves the cookie from the request header and sends it back to the
// client in the response body.
func Show(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(Name)
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, "cookie not found", http.StatusBadRequest)
		default:
			log.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}
	w.Write([]byte("found cookie" + " " + cookie.Name))
}
