// Parsejson parses a JSON string and prints it out.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

var filesystem = `
[
	{
		"path": "/etc",
		"is_dir": true
	},
	{
		"path": "/etc/hosts",
		"is_dir": false
	}
]
`

func main() {
	var fs []struct {
		Path  string
		IsDir bool `json:"is_dir"`
	}

	r := strings.NewReader(filesystem)
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&fs); err != nil {
		log.Fatalf("decoding JSON: %v", err)
	}

	for _, f := range fs {
		fmt.Println(f.Path, f.IsDir)
	}
}
