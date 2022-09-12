// Shop4 is a web shop selling shoes and socks. It registers the /list and
// /price endpoints by calling the convenience function http.HandleFunc.
// Adapted from github.com/adonovan/gopl.io/tree/master/ch7/http4.
//
// Level: intermediate
// Topics: web server, net/http, maps
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q", item)
		return
	}
	fmt.Fprintf(w, "%s", price)
}
