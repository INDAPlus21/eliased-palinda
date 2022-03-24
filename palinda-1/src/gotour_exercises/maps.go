package main

import (
	"golang.org/x/tour/wc"
	"strings"
	// "fmt"
)

func WordCount(s string) map[string]int {
	// fmt.Println(strings.Fields("  foo bar  baz   "))
	word_count := make(map[string]int)
	word_list := strings.Fields(s) 
	for _, word := range word_list {
		word_count[word] += 1
	} 
	return word_count // map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
	// fmt.Println(WordCount("hello hi hello you hi"))
	// fmt.Println(strings.Fields("  foo bar  baz   "))
}
