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

type Response struct {
	result string
}

func main() {
	tasks := make(chan Task, 1000)
	go taskProducer(tasks, 1000)

	poolSize := 5
	pool := make(chan bool, poolSize)

	// START OMIT
	responseChan := make(chan Response, 1000) // HL

	go func() {
		for {
			fmt.Println((<-responseChan).result) // HL
		}
	}()

	for {
		select {
		case task := <-tasks:
			pool <- true
			go func() {
				defer func() { <-pool }()
				time.Sleep(task.delay)
				responseChan <- Response{result: task.message} // HL
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
