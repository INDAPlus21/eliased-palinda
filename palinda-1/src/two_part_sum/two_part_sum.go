package main

import (
	"fmt"
	"sync"
)

func sumArray(array []int, sum *int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, value := range array {
		*sum += value
	}
}

func ConcurrentSum(array []int) int {
	var wg sync.WaitGroup
	sum := 0
	wg.Add(1)
	go sumArray(array[:len(array)/2], &sum, &wg)
	wg.Add(1)
	go sumArray(array[len(array)/2:], &sum, &wg)
	wg.Wait()
	fmt.Println(array, sum)
	return sum
}

func main() {
	// var array []int
	// fmt.Println("hello) 
	// var wg sync.WaitGroup
	// sum := 0
	// wg.Add(1)
		// array := []int{1, 3, 123, 23}
	//  ConcurrentSum(array)
	// wg.Add(1)
	// go ConcurrentSum(array[2:], &sum, &wg)
	// wg.Wait()
	// fmt.Println(array, sum)
}