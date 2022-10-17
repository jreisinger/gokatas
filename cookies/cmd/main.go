// Write an encrypted and authenticated cookie to each response. To see it:
//
//	curl -i localhost:3000/cookie
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

	log.Print("listening at :3000 ...")
	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal(err)
	}
}
