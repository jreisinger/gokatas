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
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
