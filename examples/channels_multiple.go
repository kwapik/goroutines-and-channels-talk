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
		var msg string
		msg = <-c1
		fmt.Println(msg)
		msg = <-c2
		fmt.Println(msg)
	}
}

// END OMIT
func cheerChannel(out chan string, msg string) {
	for i := 0; ; i++ {
		out <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
