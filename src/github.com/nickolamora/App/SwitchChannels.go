package main

import (
	"fmt"
	"time"
)

func main() {
	//create an unbuffered channel
	R1 := make(chan string)
	R2 := make(chan string)

	//generate data to a specific channel
	go portal1(R1)
	go portal2(R2)

	//With a Select Channel, we can do different scenarios if we're receiving data from different channels
	//in this case if we receive data from R1 do something, if we receive data from R2, do something else
	// to do a continuous stream, we would use this select statement in the for loop int Channels.go
	select {
	case op1 := <-R1:
		fmt.Println(op1)
	case op2 := <-R2:
		fmt.Println(op2)
	}
}

func portal2(channel2 chan string) {
	time.Sleep(1 * time.Second)
	channel2 <- "Welcome to channel 2"
}
func portal1(channel1 chan string) {
	time.Sleep(2 * time.Second)
	channel1 <- "Welcome to channel 1"
}
