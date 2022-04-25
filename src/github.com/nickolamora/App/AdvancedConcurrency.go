package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

/**
Channe and Contexts
---------------------------------

Channels are a safe way of transferring data between goroutines without using a mutex. You can send data to a channel in one goroutine, then consume data from the same channel in another goroutine. By default, a channel does not have space to store data, so you must simultaneously send and receive data from the channel to avoid a deadlock. A buffered channel allows you to allocate space to temporarily store data in the channel.

Context is an object that can be safely passed to multiple goroutines to provide a way to implement a timeout or cancellation to a function. 3rd party libraries that make HTTP requests or database requests typically have support for providing your own context so you cancel those operations.
*/
func DoWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}

func main() {
	fmt.Println("=======Channel Creation=========")

	//unbuffered
	dataChan := make(chan int)

	go func() {
		dataChan <- 123 //in channels, you need one thread publishing to a channel
	}()

	n := <-dataChan // and another thread receiving from the channel

	fmt.Printf("n = %d\n", n)

	fmt.Println("=======Channel Usage=========")

	dataChan1 := make(chan int)

	go func() {
		//tracking go routines to wait for them to complete
		wg := sync.WaitGroup{}
		for i := 0; i < 10; i++ {
			go func() {
				wg.Add(1)
				defer wg.Done()
				result := DoWork()
				dataChan1 <- result
			}()
		}
		//wait before all the iterations are done before we close the channel
		wg.Wait()
		//after we're done producing data to a channel, we have to close the channel
		close(dataChan1)
	}()

	//looping through a channel
	for n := range dataChan1 {
		fmt.Printf("n = %d\n", n)
	}

	fmt.Println("=======Context=========")
	//context is a way to add a timeout or cancellation to a goroutine.

	//create context, if it takes longer than 100 ms
	//											Here you can pass the parent context to append new things like cancellations or timeouts
	timeoutContext, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	//cancels the timeout
	defer cancel()

	//create HTTP req
	req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, "http://placehold.it/2000x2000", nil)
	if err != nil {
		panic(err)
	}

	// perform HTTP request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	// get data from HTTP response
	imageData, err := ioutil.ReadAll(res.Body)
	fmt.Printf("downloaded image of size %d\n", len(imageData))
}
