// Web server that sets a cookie (/set) and shows that you sent it back (/get).
package main

import (
	"log"
	"net/http"

	"github.com/jreisinger/gokatas/cookie"
)

func main() {
	// Start a web server with two endpoints.
	mux := http.NewServeMux()
	mux.HandleFunc("/set", cookie.Set)
	mux.HandleFunc("/get", cookie.Get)

	log.Print("listening at localhost:3000 ...")
	if err := http.ListenAndServe("localhost:3000", mux); err != nil {
		log.Fatal(err)
	}
}
