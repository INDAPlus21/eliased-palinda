package main

import "fmt"

// I want this program to print "Hello world!", but it doesn't work.
func main() {
	ch := make(chan string, 1)
	ch <- "Hello world!"
	fmt.Println(<-ch)
}

// Explanation: the channel was unbuffered, and I made it buffered a