// Stefan Nilsson 2013-03-13

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

/* 
Your program should contain two channels: One channel for questions, and one for answers and predictions. 
In the Oracle function you should start three indefinite go-routines.

A go-routine that receives all questions, and for each incoming question, creates a separate go-routine that answers that question
A go-routine that generates predictions
A go-routine that receives all answers and predictions, and prints then to stdout
*/ 

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	questions := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		questions <- line // The channel doesn't block.
	}
}

func RandomSleep(n int) {
	time.Sleep(time.Duration(rand.Intn(n)) * time.Millisecond)
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string)
	answers := make(chan string)
	// receive all questions 
	go func() {
		for question := range questions {
			fmt.Println("received", question)
			go prophecy(question, answers)
			// RandomSleep(100) // Simulate time to consume data.
		}
	}()
	go func() {
		for {
			RandomSleep(10000) // Simulate time to consume data.
			fmt.Println("This is a random prediction")
		}
	}()
	go func() {
		for answer := range answers {
			fmt.Println("received", answer)
		}
	}()
	// TODO: Answer questions.
	// TODO: Make prophecies.
	// TODO: Print answers.
	return questions
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
// TODO: make it better.
func prophecy(question string, answer chan<- string) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.
	time.Sleep(time.Duration(rand.Intn(1)) * time.Second) // deleted long wait 
	fmt.Println("in prophecy")

	// Find the longest word.
	longestWord := ""
	words := strings.Fields(question) // Fields extracts the words into a slice.
	for _, w := range words {
		if len(w) > len(longestWord) {
			longestWord = w
		}
	}

	// Cook up some pointless nonsense.
	/* 	nonsense := []string{

		// "The moon is dark.",
		// "The sun is bright.",
	} */
	answer <- longestWord + "..." + longestWord + "are indeed. " // + nonsense[rand.Intn(len(nonsense))]
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}