package main

import (
	"fmt"
	"math/rand"
	"time"
)

// UnbufferedChanBug1 close channel on receiver side
// that would get an error: `panic: send on closed channel`
// best practice, should close channel on Sender side
func UnbufferedChanBug1() {
	queue := make(chan int)

	// pull from queue
	go func() {
		for {
			select {
			case b, ok := <-queue:
				if !ok {
					fmt.Println("channel closed")
					return
				}
				fmt.Println(b)
				if b == 10 {
					fmt.Println("got exit signal, closing queue")
					close(queue)
				}
			}
		}
	}()

	// push to queue
	for {
		time.Sleep(time.Millisecond * 500)
		queue <- rand.Intn(11) // panic: send on closed channel
	}

}

// UnbufferedChanTest
// If the channel is unbuffered, the sender blocks until the receiver has received the value.
func UnbufferedChanTest() {
	fmt.Println("process start @", time.Now())
	intChan := make(chan int)

	go func() {
		for {
			val := rand.Intn(100)
			time.Sleep(time.Second * 3)
			intChan <- val
			fmt.Printf("[%v] pushing value:%d\n", time.Now(), val)
		}
	}()

	for {
		v := <-intChan // block, because nothing pushed into queue.
		fmt.Printf("[%v] received value:%d\n", time.Now(), v)
	}
}

/*
	Buffered Channel
	Q: 什麼情況下可以在長度有限的queue下塞入超過長度的message。例如：channel長度為10, 但我有100筆資料要處理 ?
	A: 使用unbuffered channel

	Q: buffered channel是blocking還是non-blocking ?
	A: 在channel未達到上限值時是後者，若已達到上限值則是前者
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
		time.Sleep(200 * time.Millisecond)
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

func BufferedChanPullBlocked() {
	intChan := make(chan int, 1)

	go func() {
		for {
			time.Sleep(time.Second * 10)
			intChan <- rand.Intn(10)
		}
	}()

	for {
		select {
		case v := <-intChan:
			fmt.Println(v) // Blocking, wait data push to channel
		}
	}
}

func BufferedChanPushBlocked() {
	intChan := make(chan int, 1)

	go func() {
		for {
			intChan <- rand.Intn(10) // Blocking, wait data pull from channel
		}
	}()

	for {
		select {
		case v := <-intChan:
			fmt.Println(v)
			time.Sleep(time.Second * 10)
		}
	}
}

func BufferedChanUnblocked() {
	intChan := make(chan int, 1)
	go func() {
		for i := 0; i < 10; i++ {
			intChan <- i
			fmt.Println("pushing", i)
			sec := rand.Intn(5)
			sp := time.Duration(sec) * (time.Second)
			time.Sleep(sp)

		}
		close(intChan)
	}()

	for {
		fmt.Println(<-intChan) //cannot stop, should check channel is closed via select case.
	}
}

func main() {
	UnbufferedChanBug1()
}
