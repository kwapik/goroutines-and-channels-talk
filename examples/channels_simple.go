// +build OMIT

package main

func main() {
	// START OMIT
	// Declaring and initializing
	c := make(chan int)

	// Sending on a channel
	c <- 42

	// Receiving from a channel
	answer := <-c
	// END OMIT
}
