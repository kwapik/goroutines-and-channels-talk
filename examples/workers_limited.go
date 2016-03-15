// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Task struct {
	message string
	delay   time.Duration
}

func main() {
	tasks := make(chan Task, 1000)
	go taskProducer(tasks, 1000)

	// START OMIT
	poolSize := 5
	pool := make(chan bool, poolSize) // HL

	for {
		select {
		case task := <-tasks:
			pool <- true // HL
			go func() {
				defer func() { <-pool }() // HL
				time.Sleep(task.delay)
				fmt.Println("Processed task:", task.message)
			}()
		default:
			fmt.Println("No tasks for now")
			time.Sleep(time.Second)
		}
	}
	// END OMIT
}

func taskProducer(out chan Task, amount int) {
	for i := 1; i <= amount; i++ {
		task := Task{
			message: fmt.Sprintf("Mambo no %d", i),
			delay:   time.Duration(rand.Intn(1e3)) * time.Millisecond,
		}
		out <- task
	}
}
