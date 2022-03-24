package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	seq := []int{0, 1} 
	return func() int {
		next := seq[len(seq)-2]+seq[len(seq)-1]
		seq = append(seq, next) 
		return next
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
