package main

import (
	"testing"
	"fmt"
	"math/rand"
	"time"
)

// test that ConcurrentSum sums an even-length array correctly
func TestSumConcurrentCorrectlySumsEvenArray(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := 55

	actual := ConcurrentSum(arr)

	if actual != expected {
		t.Errorf("expected %d, was %d", expected, actual)
	}
}

func stableSum(arr []int) int {
	sum := 0 
	for _, i := range arr {
		fmt.Println(i)
		sum += i 
	}
	return sum 
}

// i'm not following the java test standards in go! 
func TestOdd(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	expected := stableSum(arr)
	actual := ConcurrentSum(arr)

	if actual != expected {
		t.Errorf("expected %d, was %d", expected, actual)
	}
}

// i'm not following the java test standards in go! 
func TestBig(t *testing.T) {
	rand.Seed(time.Now().Unix())
	arr := rand.Perm(10)

	expected := stableSum(arr)
	actual := ConcurrentSum(arr)

	if actual != expected {
		t.Errorf("expected %d, was %d", expected, actual)
	}
} 