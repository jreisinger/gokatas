// Areader reads and prints three bytes from Areader.
package main

import (
	"fmt"

	"github.com/jreisinger/gokatas/areader"
)

func main(){
	var r areader.Areader
	slice := make([]byte, 34)
	r.Read(slice)
	fmt.Println(string(slice))
}

//create a var of type Areader,
//create a slice of byte, and print it
