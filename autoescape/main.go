// Autoescape shows how to suppress html/template auto-escaping behavior for
// fields that contain trusted HTML data. (The same can be done for trusted
// JavaScript, CSS and URLs.) Taken from gopl.io ch4.6 Text and HTML templates.
//
// Level: beginner
// Topics: html/template
package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("autoescape").Parse(templ))
	var data struct {
		A string        // untrusted plain text
		B template.HTML // trusted HTML
	}
	data.A = "<b>hello</b>"
	data.B = "<b>hello</b>"
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
