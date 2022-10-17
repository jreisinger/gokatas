// Web server that sends you a cookie (/set) and shows that you sent it back
// (/get).
package main

import (
	"log"
	"net/http"

	"github.com/jreisinger/gokatas/cookies"
)

func main() {
	// Start a web server with two endpoints.
	mux := http.NewServeMux()
	mux.HandleFunc("/set", cookies.Set)
	mux.HandleFunc("/get", cookies.Get)

	log.Print("listening at localhost:3000 ...")
	if err := http.ListenAndServe("localhost:3000", mux); err != nil {
		log.Fatal(err)
	}
}
