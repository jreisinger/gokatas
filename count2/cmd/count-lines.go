package main

import (
	"fmt"
	"log"

	"github.com/jreisinger/gokatas/count2"
)

func main() {
	c, err := count2.NewCounter()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c.Lines())
}
