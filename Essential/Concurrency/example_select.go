package main

import (
	"fmt"
	"time"
)

func DoSomething1(channel chan string) {

	// wait for 3 second
	time.Sleep(3 * time.Second)

	// send data to channel
	channel <- "do something 1"
}

var start time.Time

func init() {
	start = time.Now()
}

func DoSomething2(channel chan string) {

	// wait for 5 second
	time.Sleep(5 * time.Second)

	// send data to channel
	channel <- "do something 2"
}

func main() {

	// define new channel1
	channel1 := make(chan string)
	// define new channel2
	channel2 := make(chan string)

	go DoSomething1(channel1)
	go DoSomething2(channel2)

	// blocking select, will wait until channel ready to receive data
	select {
	case result1 := <-channel1:
		fmt.Println(result1)
		fmt.Println("response time ", time.Since(start))
	case result2 := <-channel2:
		fmt.Println(result2)
		fmt.Println("response time ", time.Since(start))
	}
}
