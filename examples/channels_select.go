// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() { cheerChannel(c1, "Heeeey!") }()
	go func() { cheerChannel(c2, "Hello!") }()

	for {
		select {
		case msg := <-c1:
			fmt.Println(msg)
		case msg := <-c2:
			fmt.Println(msg)
		default:
		}
	}
}

// END OMIT
func cheerChannel(out chan string, msg string) {
	for i := 0; i < 5; i++ {
		out <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
