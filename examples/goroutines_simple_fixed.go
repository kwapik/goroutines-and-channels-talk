// +build OMIT

package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	go cheer() // HL
	time.Sleep(time.Second)
}

func cheer() {
	for i := 0; ; i++ {
		fmt.Println("Yay! I'm a goroutine!", i)
		time.Sleep(time.Second)
	}
}

// END OMIT
