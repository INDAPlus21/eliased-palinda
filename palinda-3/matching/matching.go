// http://www.nada.kth.se/~snilsson/concurrency/
package main

import (
	"fmt"
	"sync"
	"log"
)

// This programs demonstrates how a channel can be used for sending and
// receiving by any number of goroutines. It also shows how  the select
// statement can be used to choose one out of several communications.
func main() {
	people := []string{"Anna", "Bob", "Cody", "Dave", "Eva"}
	match := make(chan string, 1) // Make room for one unmatched send.
	// var wg sync.WaitGroup
	wg := new(sync.WaitGroup)
	wg.Add(len(people))
	for _, name := range people {
		log.Println("in range loop")
		go Seek(name, match, wg)
	}
	wg.Wait()
	log.Println("after waitgroup")
	select {
	case name := <-match:
		fmt.Printf("No one received %sâ€™s message.\n", name)
	default:
		log.Println("in default")
		// There was no pending send operation.
	}
}

// Seek either sends or receives, whichever possible, a name on the match
// channel and notifies the wait group when done.
func Seek(name string, match chan string, wg *sync.WaitGroup) {
	log.Println("in seek: ", name, wg)
	select {
	case peer := <-match:
		fmt.Printf("%s sent a message to %s.\n", peer, name)
	case match <- name:
		// Wait for someone to receive my message.
	}
	log.Println("ending seek, waitgroup: ", name, wg)
	wg.Done()
}

