package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// This program should go to 11, but it seemingly only prints 1 to 10.
func main() {
	ch := make(chan int)

	wg.Add(2)
	go Print(ch)
	go func() {
		defer wg.Done() // 3
		for i := 1; i <= 11; i++ {
			ch <- i // here it sends the numbers to the channel
		}
		close(ch)
	}()

	wg.Wait()
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch chan int) {
	defer wg.Done()
	for n := range ch { // reads from channel until it's closed
		time.Sleep(10 * time.Millisecond) // simulate processing time
		fmt.Println(n)
	}
}
