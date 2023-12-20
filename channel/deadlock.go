package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	- unbuffered channel is blocking method
	- buffered channel is non-blocking method
*/

func DeadLock1() {
	queue := make(chan int, 10)
	for i := 0; i < 20; i++ {
		queue <- rand.Intn(100) // Deadlock! because the queue have overflow
	}
	fmt.Println("end of function DeadLock1")
}

func Solution1() {
	queue := make(chan int, 10)
	go func() {
		for i := 0; i < 20; i++ {
			queue <- rand.Intn(100) // would block when i=10
		}
	}()
	fmt.Println("end of function Solution1")
}

func DeadLock2() {
	queue := make(chan int)
	queue <- 1

	result := <-queue * 3 * 4 // Deadlock!
	// because there are 2 actions push & pull run at single routine, they would block each other,
	// we should run push action in a goroutine
	fmt.Println("result:", result)
}

func Solution2() {
	queue := make(chan int)
	go func() {
		queue <- 1
	}()

	result := <-queue * 3 * 4
	fmt.Println("result:", result)
}

func Solution21() {
	queue := make(chan int, 1)
	queue <- 1

	result := <-queue * 3 * 4
	fmt.Println("result:", result)
}

func Test1() {
	queue := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			queue <- i
		}
	}()
	fmt.Println("--------middle gap --------")

	for i := 0; i < 10; i++ {
		select {
		case v := <-queue:
			fmt.Println("got value:", v)
		}
	}
}

func main() {
	Test1()
}
