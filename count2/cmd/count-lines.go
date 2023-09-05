package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jreisinger/gokatas/count2"
)

func main() {
	c, err := count2.NewCounter(count2.WithInputFromArgs(os.Args[1:]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c.Lines())
}
