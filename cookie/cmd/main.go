// Web server that sets a cookie (/set) and shows that you sent it back (/get).
package main

import (
	"log"
	"net/http"

	"github.com/jreisinger/gokatas/cookie"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/set", cookie.Set)
	mux.HandleFunc("/get", cookie.Get)
	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal(err)
	}
}
