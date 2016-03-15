// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT

type Task struct {
	message string
	delay   time.Duration
}

func main() {
	tasks := make(chan Task, 1000)
	go taskProducer(tasks, 1000)

	for {
		select {
		case task := <-tasks:
			go func() {
				time.Sleep(task.delay)
				fmt.Println("Processed task:", task.message) // HL
			}()
		default:
			fmt.Println("No tasks for now")
			time.Sleep(time.Second)
		}
	}
}

// END OMIT
func taskProducer(out chan Task, amount int) {
	for i := 1; i <= amount; i++ {
		task := Task{
			message: fmt.Sprintf("Mambo no %d", i),
			delay:   time.Duration(rand.Intn(1e3)) * time.Millisecond,
		}
		out <- task
	}
}
