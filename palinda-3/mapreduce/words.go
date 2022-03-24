package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"
	"time"
)

const DataFile = "loremipsum.txt"

func mäp(text []string, freq_chan chan map[string]int, wg *sync.WaitGroup) {
	freqs := make(map[string]int)
	for _, word := range text {
		freqs[word] += 1
	}
	freq_chan <- freqs
	wg.Done()
}

func reduce(freq_chan <-chan map[string]int) map[string]int {
	freqs := make(map[string]int)
	for one_map := range freq_chan {
		for k, v := range one_map {
			freqs[k] += v
		}
	}
	return freqs
}

// Return the word frequencies of the text argument.
//
// Split load optimally across processor cores.
func WordCount(text string) map[string]int {

	word_array := strings.Fields(text)
 
	for i, word := range word_array {
		word = strings.ToLower(word)
		word = strings.ReplaceAll(word, ".", "")
		word_array[i] = strings.ReplaceAll(word, ",", "")
	}

	num_threads := 13 // was optimal (through running it multiple times)
	part_size := len(word_array)/num_threads 
	freq_chan := make(chan map[string]int, num_threads+1) // because the channel needs such a large buffer
 
	wg := new(sync.WaitGroup)
 
	for i := 0; i*part_size < len(word_array); i++ {
		wg.Add(1)
		var text_part []string
		if (i+1)*part_size < len(word_array) {
			text_part = word_array[i*part_size : (i+1)*part_size]
		} else {
			text_part = word_array[i*part_size:]
		}
		go mäp(text_part, freq_chan, wg)
	}

	wg.Wait() // if you don't wait it gets closed immediately 
	close(freq_chan) // if you don't close you get deadlock

	// reducers can't start while mappers are still ongoing 
	freqs := reduce(freq_chan)

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
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	data, err := ioutil.ReadFile(DataFile)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("word count: ", WordCount(string(data)))

	numRuns := 100
	runtimeMillis := benchmark(string(data), numRuns)
	printResults(runtimeMillis, numRuns)
}
