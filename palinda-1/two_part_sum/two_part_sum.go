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
	sum := 0

	var wg sync.WaitGroup
	wg.Add(2)
	go sumArray(array[:len(array)/2], &sum, &wg)
	go sumArray(array[len(array)/2:], &sum, &wg)
	wg.Wait()
	
	fmt.Println(array, sum)
	return sum
}

func main() {
	array := []int{1, 3, 123, 23}
	ConcurrentSum(array)
}