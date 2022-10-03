// Write an encrypted and authenticated cookie to each response. To see it:
//
//	curl -i localhost:3000/cookie
package main

import (
	"encoding/hex"
	"log"
	"net/http"

	"github.com/jreisinger/gokatas/cookies"
)

// In real app this should be read at runtime from and environment variable.
const randString = "13d6b4dff8f84a10851021ec8608f814570d562c92fe6b5ec4c9f595bcb3234b"

var secretKey []byte

func main() {
	var err error
	secretKey, err = hex.DecodeString(randString)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/cookie", setCookieHandler)

	log.Print("Listening at :3000 ...")
	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal(err)
	}
}

func setCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "exampleCookie",
		Value:    "Hello Zoë!",
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	err := cookies.WriteEncrypted(w, cookie, secretKey)
	if err != nil {
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("cookie set!"))
}
