package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fibo_numbers := []int{0, 1} 
	return func() int {
		new_number := fibo_numbers[len(fibo_numbers)-2:][0]+fibo_numbers[len(fibo_numbers)-2:][1]
		fibo_numbers = append(fibo_numbers, new_number) 
		return new_number
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	// var fibo_numbers[]int 
	// fibo_numbers = append(fibo_numbers, 0)
	// fibo_numbers = append(fibo_numbers, 1)7

	// fmt.Println(fibo_numbers)
}
