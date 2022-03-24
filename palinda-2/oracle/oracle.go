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
			// fmt.Println("received", question)
			go prophecy(question, answers)
			// RandomSleep(100) // Simulate time to consume data.
		}
	}()
	go func() {
		for {
			RandomSleep(10000) // Simulate time to consume data.
			switch rand.Intn(3) {
			case 0:
				answers <- "The one with the power to vanquish the Dark Lord approaches... Born to those who have thrice defied him, born as the seventh month dies... and the Dark Lord will mark him as his equal, but he will have power the Dark Lord knows not... and either must die at the hand of the other for neither can live while the other survives... The one with the power to vanquish the Dark Lord will be born as the seventh month dies..."
			case 1:
				answers <- "It will happen tonight. The Dark Lord lies alone and friendless, abandoned by his followers. His servant has been chained these twelve years. Tonight, before midnight... the servant will break free and set out to rejoin his master. The Dark Lord will rise again with his servant's aid, greater and more terrible than ever he was. Tonight... before midnight... the servant... will set out... to rejoin... his master..."
			case 2:
				answers <- "The Grim, my dear, the Grim!’ cried Professor Trelawney, who looked shocked that Harry hadn’t understood. ‘The giant, spectral dog that haunts churchyards! My dear boy, it is an omen – the worst omen – of death!"
			}
		}
	}()
	go func() {
		for answer := range answers {
			fmt.Println(answer)
		}
	}()
	// TODO: Answer questions.
	// TODO: Make prophecies.
	// TODO: Print answers.
	return questions
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
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

	answer <- (longestWord + "... " + longestWord + " indeed...") // + nonsense[rand.Intn(len(nonsense))]
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}