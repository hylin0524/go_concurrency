package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	To prevent our server or remote server explosion (too many threads)
	we can proactively limit the number of goroutines by using a goroutine pool.
*/

const maxGoroutine = 10

func printer(i, worker int) {
	r := i % 10
	time.Sleep(time.Second * time.Duration(r))                                                        // mock api server response time
	fmt.Printf("processing number [number %d] on [worker %d] at [time %v] \n", i, worker, time.Now()) // mock api processed result
}

func ChanWithWaitGroup() {
	intChan := make(chan int, 20)
	var wg sync.WaitGroup

	wg.Add(maxGoroutine) // limit goroutines
	for i := 0; i < maxGoroutine; i++ {
		go func(w int) {
			defer wg.Done()
			// infinite loop and pull value from channel every goroutine
			for {
				v, ok := <-intChan
				if !ok { // if there is no job then exit the loop
					return
				}
				printer(v, w)
			}
		}(i)
	}

	for j := 0; j < 100; j++ { // mock total jobs
		intChan <- j
	}

	close(intChan)
	wg.Wait()
}

func main() {
	ChanWithWaitGroup()
}
