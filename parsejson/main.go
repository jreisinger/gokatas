// Parsejson parses a JSON string and prints it out.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

var filesystem = `[
	{
		"path": "/etc",
		"is_dir": true
	},
	{
		"path": "/etc/hosts",
		"is_dir": false
	}
]`

func main() {
	r := strings.NewReader(filesystem)

	var fs []struct {
		Path  string
		IsDir bool `json:"is_dir"`
	}

	if err := json.NewDecoder(r).Decode(&fs); err != nil {
		log.Fatalf("decoding JSON: %v", err)
	}

	fmt.Printf("%-10v\t%v\n", "Path", "Dir")
	fmt.Printf("%-10v\t%v\n", "----", "---")
	for _, f := range fs {
		fmt.Printf("%-10s\t%t\n", f.Path, f.IsDir)
	}
}
