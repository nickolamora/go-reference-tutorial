package main

import (
	"fmt"
	"sync"
)

//Channels - medium of data flow for goroutines
// generally sending and receiving data in and out of goroutines is accomplished using channels

var wg2 = sync.WaitGroup{}

func Main() {
	//creating a channel
	//buffered channels - hold all the data that is available to avoid deadlock!
	ch := make(chan int, 100)

	//create two goroutine functions to send and receive data to and from this channel
	//pass in how much data it can hold and loop through it all
	wg2.Add(2)
	go func(ch <-chan int) {
		//receiving data from a channel
		for {
			//check if there is still data coming in using the 'ok' from the channel, if not close out.
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg2.Done()
	}(ch)
	go func(ch chan<- int) {
		// sending data to a channel
		ch <- 12
		//close the channel to tell the consuming channel you're done producing
		close(ch)
		wg2.Done()
	}(ch)
	wg2.Wait()
	//bi directional channels (not recommended in prod)

}
