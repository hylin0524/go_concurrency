package main

import (
	"fmt"
	"math/rand"
	"time"
)

func UnbufferedChan() {
	queue := make(chan int)

	// push to queue
	go func() {
		for {
			time.Sleep(time.Millisecond * 200)
			queue <- rand.Intn(11)
		}
	}()

	// pull from queue
	for {
		select {
		case b := <-queue:
			fmt.Println(b)
			if b == 10 {
				fmt.Println("got exit signal")
				return
			}
		}
	}
}

/*
	Buffered Channel
	Q1: 什麼情況下可以在長度有限的queue下塞入超過長度的message。例如：channel長度為10, 但我有100筆資料要處理 ?
*/

func BufferedChan() {
	queue := make(chan int, 5)

	go func() {
		for i := 0; i < 50; i++ {
			queue <- i
		}
		close(queue)
	}()

	for {
		time.Sleep(100 * time.Millisecond)
		select {
		case v, ok := <-queue:
			if !ok { // the condition of exit infinite-loop
				fmt.Println("Channel closed")
				return
			}
			fmt.Println("got value:", v)
		default:
			fmt.Println("nothing to do, waiting...")
		}
	}
}

func main() {

}
