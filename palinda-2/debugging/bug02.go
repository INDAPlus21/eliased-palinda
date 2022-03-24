package main

import (
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup

// This program should go to 11, but it seemingly only prints 1 to 10.
func main() {
	ch := make(chan int)
	wg.Add(1)
	go Print(ch)
	// The error is caused by it not WAITING for the goroutine to finish 
	wg.Add(1)
	go func() {
		defer wg.Done() // 3
		for i := 1; i <= 11; i++ {
		ch <- i // here it sends the numbers to the channel 
	}
	}()
	// so it's deadlock because receiving/sending to a channel is a blocking operation (???)4
	// until "the other side" is done, buffered wait until the whole capacity is full 
	// close = not able to receive or send any more values 
	// range pulls values "automatically" until channel is closed 
	// select acts like a switch 
	// So the reason for the bug is that the channel is immediately closed after the 11 has been added 
	// and then because Print() sleeps, it doesn't have time to print it 
	wg.Wait() // Wait for all producers to finish.
	close(ch)
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int) {
	defer wg.Done() // 3
	for n := range ch { // reads from channel until it's closed
		time.Sleep(10 * time.Millisecond) // simulate processing time
		fmt.Println(n)
	}
	// close(ch)
}