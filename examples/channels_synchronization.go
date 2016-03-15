// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT
func main() {
	c := make(chan int)

	go func() {
		cheer("Yay! I'm a goroutine!")
		c <- 1 // HL
	}()

	<-c // HL
	fmt.Println("Bye!")
}

// END OMIT

func cheer(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
