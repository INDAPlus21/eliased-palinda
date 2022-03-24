package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	word_count := make(map[string]int)
	word_list := strings.Fields(s) 
	for _, word := range word_list {
		word_count[word] += 1
	} 
	return word_count 
}

func main() {
	wc.Test(WordCount)
}
