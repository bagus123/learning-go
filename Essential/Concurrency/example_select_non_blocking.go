package main

import (
	"fmt"
	"time"
)

func DoSomething1(channel chan string) {
	channel <- "do something 1"
}

var start time.Time

func init() {
	start = time.Now()
}

func DoSomething2(channel chan string) {
	channel <- "do something 2"
}

func main() {

	channel1 := make(chan string)
	channel2 := make(chan string)

	go DoSomething1(channel1)
	go DoSomething2(channel2)

	// waiting for 3 second
	// if this code line comment, will execute default case
	time.Sleep(3 * time.Second)

	select {
	case result1 := <-channel1:
		fmt.Println(result1)
		fmt.Println("response time ", time.Since(start))
	case result2 := <-channel2:
		fmt.Println(result2)
		fmt.Println("response time ", time.Since(start))
	default:
		fmt.Println("No goroutine available to send data")

	}
}
