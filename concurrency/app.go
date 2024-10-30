package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)

	doneChan <- true // sends data to channel
}

func main() {
	// [go keyword]:
	// adding go keyword in front of function invocation
	// tells go to run these function as "go routines" (parallel and run the function in a non-blocking way)

	// however, adding go keyword makes function dispatch for goroutines and does not wait for to print (fmt.Println())
	// in command line before it exits program

	// go greet("Nice to meet you!")
	// go greet("How are you?")

	// [channel]:
	// a value that can transmit data and be used as a communication tool when working with go routine

	done := make(chan bool)
	go slowGreet("How ... are ... you ...?", done)

	<-done // this works as go to wait for exit the program until channel receives the data

	// go greet("I hope you're liking the course!")
}
