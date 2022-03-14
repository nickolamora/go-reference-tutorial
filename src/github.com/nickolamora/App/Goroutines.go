package main

import (
	"fmt"
	"sync"
	"time"
)

// Goroutines - a lightweight thread in GO.
// multiple 'goroutines' can be forked off of the main Goroutine (thread) but if the main one finishes, they all finish.

var wg = sync.WaitGroup{}

func main() {
	//executing a GO routine
	go writeGoMessage()
	fmt.Println("================")

	//example of a race condition since the 'message' variable can change faster than the goroutine can finish execution
	message := "Hello User"
	go func() {
		fmt.Println(message)
	}()
	message = "Hello Youtube"
	time.Sleep(11 * time.Millisecond)

	fmt.Println("================")

	// how to avoid a race condition
	// the best way to avoid it is to pass the variable within the function
	// you cand debug if you have a race condition if you run your program using 'go run -race Main.go'
	message2 := "Hello User"
	go func(message2 string) {
		fmt.Println(message2)
	}(message2)
	message2 = "Hello Youtube"
	fmt.Println(message2)
	time.Sleep(11 * time.Millisecond)

	fmt.Println("================")
	//How to wait for a goroutine to finish before ending the Main goroutine
	message3 := "Hello User"
	//before starting the goroutine, tell Go that there is one goroutine to wait for by adding it to the waitGroup (go routines to wait for before closing main goroutine)
	wg.Add(1)
	go func(message string) {
		fmt.Println(message3)
		// tell the compiler your functionality is done by adding 'Done' to when you finish
		wg.Done()
	}(message3)
	message3 = "Hello Youtube"
	// at the end of the method tell go to wait
	wg.Wait()
	time.Sleep(11 * time.Millisecond)
}

func writeGoMessage() {
	fmt.Println("Hello")
}
