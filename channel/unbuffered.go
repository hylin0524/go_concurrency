package main

import (
	"fmt"
	"math/rand"
	"time"
)

// UnbufferedChanBug1 close channel on receiver side
// that would get an error: `panic: send on closed channel`
// BEST PRACTICE: close channel on sender side
func UnbufferedChanBug1() {
	queue := make(chan int)

	// push to queue
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			queue <- rand.Intn(11) // panic: send on closed channel
		}
	}()

	// pull from queue
	for {
		select {
		case b, ok := <-queue:
			if !ok {
				fmt.Println("channel closed")
				return
			}
			fmt.Println("pulling data:", b)
			if b == 10 {
				fmt.Println("got exit signal, closing queue")
				close(queue)
			}
		}
	}
}

// UnbufferedChan
// If the channel is unbuffered, receiver would be blocked when no value in channel.
func UnbufferedChan() {
	fmt.Println("process start @", time.Now())
	intChan := make(chan int)

	go func() {
		for {
			val := rand.Intn(100)
			time.Sleep(time.Second * 3)
			fmt.Printf("[%v] pushing value:%d\n", time.Now(), val)
			intChan <- val
			fmt.Printf("[%v] pushed value:%d\n", time.Now(), val)
		}
	}()

	for {
		v := <-intChan // block, because nothing pushed into queue.
		fmt.Printf("[%v] received value:%d\n", time.Now().String(), v)
	}
}

func Poc() {
	queue := make(chan int)
	queue <- 1
	result := <-queue + 100
	fmt.Println(result)
}

func main() {
	Poc()
}
