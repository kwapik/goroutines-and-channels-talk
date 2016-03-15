// +build OMIT

package main

func main() {
	// START OMIT
	c := make(chan int, 2)

	c <- 13
	c <- 37

	first := <-c
	second := <-c
	// END OMIT
}
