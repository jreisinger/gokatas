package main

import (
	"log"
	"regexp"

	"github.com/jreisinger/gokatas/match"
)

func main() {
	m, err := match.NewMatcher()
	if err != nil {
		log.Fatal(err)
	}
	m.PrintLines(regexp.MustCompile("hello"))
}
