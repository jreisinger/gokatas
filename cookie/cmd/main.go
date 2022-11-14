// Web server that sets a cookie and shows that client sent it back.
package main

import (
	"log"
	"net/http"

	"github.com/jreisinger/gokatas/cookie"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/set", cookie.Set)
	mux.HandleFunc("/show", cookie.Show)
	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal(err)
	}
}
