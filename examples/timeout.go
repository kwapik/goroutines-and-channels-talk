// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT
func main() {
	c := make(chan string)

	go func() { cheerChannel(c, "Heeeey!") }()

	for {
		select {
		case msg := <-c:
			fmt.Println(msg)
		case <-time.After(time.Second): // HL
			fmt.Println("Timeout!")
			return
		}
	}
}

// END OMIT
func cheerChannel(out chan string, msg string) {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; ; i++ {
		out <- fmt.Sprintf("%s %d", msg, i)
		t := rand.Intn(1500)
		time.Sleep(time.Duration(t) * time.Millisecond)
	}
}
