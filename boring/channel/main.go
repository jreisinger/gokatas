// A channel allows for communication and synchronization between goroutines.
//
// Level: beginner
// Topics: goroutines, channels
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c:=make(chan string)
	go boring("oi", c)
	for i:=0; i <7;i++{
		fmt.Println(<- c)
	}
}

func boring(msg string, c chan string){
	for i :=0;;i++{
		c <- fmt.Sprintf("%s %d \n", msg, i)
		tmp:=rand.Intn(1e3)
		time.Sleep(time.Duration(tmp)*time.Millisecond)
	}
}

//criar uma func boring que recebe uma str e um chan de param
//chamar boring na main
