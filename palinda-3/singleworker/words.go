package main

import (
	"fmt"
	"time"
	"log"
	// "os"
	"io/ioutil"
	"strings"
)

const DataFile = "loremipsum.txt"

// Return the word frequencies of the text argument.
func WordCount(text string) map[string]int {
	freqs := make(map[string]int)
	words := strings.Fields(text)
	for _, word := range words {
		formatted_word := strings.ToLower(strings.ReplaceAll(word, ".", ""))
		formatted_word = strings.ReplaceAll(formatted_word, ",", "")
		freqs[formatted_word] += 1
	}

	/* for word := range(text) {
		freqs[word] += 1
	} */
	// log.Println(freqs)
	// ...
	return freqs
}

// Benchmark how long it takes to count word frequencies in text numRuns times.
//
// Return the total time elapsed.
func benchmark(text string, numRuns int) int64 {
	start := time.Now()
	for i := 0; i < numRuns; i++ {
		WordCount(text)
	}
	runtimeMillis := time.Since(start).Nanoseconds() / 1e6

	return runtimeMillis
}

// Print the results of a benchmark
func printResults(runtimeMillis int64, numRuns int) {
	fmt.Printf("amount of runs: %d\n", numRuns)
	fmt.Printf("total time: %d ms\n", runtimeMillis)
	average := float64(runtimeMillis) / float64(numRuns)
	fmt.Printf("average time/run: %.2f ms\n", average)
}

func main() {
	// read in DataFile as a string called data
	data, err := ioutil.ReadFile(DataFile)
	if err != nil {
		log.Println(err.Error())
	}
	/* data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err.Error())
	} */

	log.Println(string(data))

	log.Println("word count: ", WordCount(string(data)))

	numRuns := 100
	runtimeMillis := benchmark(string(data), numRuns)
	printResults(runtimeMillis, numRuns)
}