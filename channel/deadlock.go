package main

import (
	"fmt"
	"math/rand"
)

/*
	- unbuffered channel is blocking method
	- buffered channel is non-blocking
		- when pulling empty chan -> blocking
		- when pushing full chan-> blocking
*/

func BufferedChanDeadLock() {
	queue := make(chan int, 10)
	for i := 0; i < 20; i++ {
		queue <- rand.Intn(100) // Deadlock!
		// because the queue have full, cannot push more message.
		// and there are no receiver pull from queue.
	}
	fmt.Println("end of function DeadLock1")
}

func BufferedChanSolution1() {
	queue := make(chan int, 10)
	go func() {
		for i := 0; i < 50; i++ {
			queue <- i // would block when i=10
		}
		close(queue)
	}()

	for {
		select {
		case v, ok := <-queue:
			if !ok {
				fmt.Println("channel closed")
				return
			}
			fmt.Println("got value:", v)
		default:
			fmt.Println("waiting ...")
		}
	}
}

func UnbufferedChanDeadLock() {
	queue := make(chan int)
	queue <- 1

	result := <-queue * 3 * 4 // Deadlock!
	// because unbuffered queue is blocking
	// there are 2 actions push & pull run at single routine, they would block each other, we should run push action in a goroutine
	fmt.Println("result:", result)
}

// UnbufferedChanSolution
// avoid 2 blocking process run on same goroutine
func UnbufferedChanSolution() {
	queue := make(chan int)
	go func() {
		queue <- 1
	}()

	result := <-queue * 3 * 4
	fmt.Println("result:", result)
}

// UnbufferedChanSolution2
// make channel as non-blocking (buffered is non-blocking)
func UnbufferedChanSolution2() {
	queue := make(chan int, 1)
	queue <- 1
	result := <-queue * 3 * 4
	fmt.Println("result:", result)
}

func main() {
	BufferedChanDeadLock()
}
