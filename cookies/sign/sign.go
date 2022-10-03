// Package sign ensures integrity of cookies, i.e. cookies are tamper-proof.
//
// Level: intermediate
// Topics: signing
package sign

import (
	"crypto/hmac"
	"crypto/sha256"
	"net/http"

	"github.com/jreisinger/gokatas/cookies"
)

func Write(w http.ResponseWriter, cookie http.Cookie, secretKey []byte) error {
	mac := hmac.New(sha256.New, secretKey)
	mac.Write([]byte(cookie.Name))
	mac.Write([]byte(cookie.Value))
	signature := mac.Sum(nil)
	cookie.Value = string(signature) + cookie.Value
	return cookies.Write(w, cookie)
}
