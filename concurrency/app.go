package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)

	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)

	doneChan <- true // sends data to channel
	// closing channel. channel can not longer receive values in other functions as well.
	// without adding close() to the function that has longest operation,
	//
	close(doneChan)
}

func main() {
	// [go keyword]:
	// adding go keyword in front of function invocation
	// tells go to run these function as "go routines" (parallel and run the function in a non-blocking way)
	// Function with go keyword won't be able to return value. If we want to return value, we have to use channel (check price-calculator-2 for reference)

	// however, adding go keyword makes function dispatch for goroutines and does not wait for to print (fmt.Println())
	// in command line before it exits program

	// [channel]:
	// a value that can transmit data and be used as a communication tool when working with go routine
	// it can receive single/multiple values

	done := make(chan bool)

	go greet("Nice to meet you!", done)
	go greet("How are you?", done)

	go slowGreet("How ... are ... you ...?", done)

	go greet("I hope you're liking the course!", done)

	// this works as go to wait for exit the program until channel receives the data
	// <-done

	// but if we want to wait for the multiple operations, it is better to use loop
	for range done {
		// we don't need body if we only want to wait for the all functions executions are completed
		// channel loop is written on price-calculator-2 as well for the reference
	}

}
