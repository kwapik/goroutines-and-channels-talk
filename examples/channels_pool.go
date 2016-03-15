// +build OMIT

package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	c := make(chan int)

	go func() {
		fmt.Println("Yay! I'm a goroutine!")
		time.Sleep(time.Second)
		c <- 1
	}()

	<-c
	fmt.Println("Bye!")
}

// END OMIT
