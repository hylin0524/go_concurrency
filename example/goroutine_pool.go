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

func printer(i int) {
	fmt.Println("handling with number:", i)
	time.Sleep(time.Second) // mock api server response time
}

func main() {
	intChan := make(chan int)
	var wg sync.WaitGroup

	wg.Add(maxGoroutine) // limit goroutines
	for i := 0; i < maxGoroutine; i++ {
		go func() {
			defer wg.Done()
			// infinite loop and pull value from channel every goroutine
			for {
				v, ok := <-intChan
				if !ok { // if there is no job then exit the loop
					return
				}
				printer(v)
			}
		}()
	}

	for j := 0; j < 100; j++ { // mock total jobs
		intChan <- j
	}

	close(intChan)
	wg.Wait()
}
