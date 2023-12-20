package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	queue := make(chan int)

	// push to queue
	go func() {
		for {
			time.Sleep(time.Second * 3)
			queue <- rand.Intn(11)
			time.Sleep(time.Second)
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
