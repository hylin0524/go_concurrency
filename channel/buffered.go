package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Buffered Channel
	Q: 什麼情況下可以在長度有限的queue下塞入超過長度的message。例如：channel長度為10, 但我有100筆資料要處理 ?
	A: 使用unbuffered channel

	Q: buffered channel是blocking還是non-blocking ?
	A: 在channel未達到上限值時是non-blocking，若已達到上限值則是blocking
*/

func BufferedChan() {
	queue := make(chan int, 5)

	go func() {
		for i := 1; i <= 50; i++ {
			queue <- i
		}
		close(queue)
	}()

	for {
		time.Sleep(200 * time.Millisecond) // pull msg every 0.2 second
		select {
		case v, ok := <-queue:
			if !ok { // the condition of exit infinite-loop
				fmt.Println("channel closed")
				return
			}
			fmt.Println("got value:", v)
		default:
			fmt.Println("nothing to do, waiting...")
		}
	}
}

func BufferedChanPullBlocked() {
	intChan := make(chan int, 1)

	go func() {
		for {
			time.Sleep(time.Second * 3)
			value := rand.Intn(10)
			fmt.Println("inserting data:", value)
			intChan <- value
		}
	}()

	for {
		select {
		case v := <-intChan:
			fmt.Println("pulling data:", v)
			fmt.Println(v) // Blocking, wait data push to channel
		}
	}
}

func BufferedChanPushBlocked() {
	intChan := make(chan int, 1)

	go func() {
		for {
			value := rand.Intn(10)
			intChan <- value // Blocking, wait data pull from channel
			fmt.Println("pushing data:", value)
		}
	}()

	for {
		time.Sleep(time.Second * 3)
		select {
		case v := <-intChan:
			fmt.Println("pulling data:", v)
		}
	}
}

func BufferedChanUnblocked() {
	intChan := make(chan int, 1)
	go func() {
		for i := 0; i < 10; i++ {
			intChan <- i
			fmt.Println("pushing data", i)
			time.Sleep(100 * time.Millisecond)
		}
		close(intChan)
	}()

	for {
		fmt.Println(<-intChan) // cannot stop, should check channel is closed via select case.
	}
}

func main() {
	BufferedChanUnblocked()
}
