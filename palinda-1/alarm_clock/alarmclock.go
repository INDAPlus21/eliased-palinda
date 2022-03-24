package main

import (
	"fmt"
	"time"
)

func Remind(text string, delay time.Duration) { 
	for {
		now := time.Now()
		fmt.Printf("The time is %s: %s \n", now.Format("15.04.05"), text) 
		time.Sleep(delay)
	}
}

func main() {
	go Remind("Time to eat", 10*time.Second)
	go Remind("Time to work", 30*time.Second)
	Remind("Time to sleep", 60*time.Second)
}

// run with go test *.go 