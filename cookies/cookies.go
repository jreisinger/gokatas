// Package cookies writes, signes and encrypts cookies [1]. An HTTP cookie is a
// small piece of data that a server sends to a user's web browser. Cookies are
// used mainly for:
//
//   - Session management (e.g. logins, shopping carts)
//   - Personalization (e.g. user preferences, themes)
//   - Tracking (recording and analyzing user behavior)
//
// Based on https://www.alexedwards.net/blog/working-with-cookies-in-go
//
// [1] https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies
//
// Level: intermediate
// Topics: signing, ecnryption, security
package cookies

import (
	"encoding/base64"
	"errors"
	"net/http"
)

var (
	ErrValueTooLong = errors.New("cookie value too long")
	ErrInvalidValue = errors.New("invalid cookie value")
)

func Write(w http.ResponseWriter, cookie http.Cookie) error {
	cookie.Value = base64.URLEncoding.EncodeToString([]byte(cookie.Value))
	if len(cookie.String()) > 4096 {
		return ErrValueTooLong
	}
	http.SetCookie(w, &cookie)
	return nil
}
