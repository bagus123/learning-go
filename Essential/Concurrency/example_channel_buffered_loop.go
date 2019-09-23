package main

import (
	"fmt"
	"time"
)

func broadcast(channel chan int) {
	for i := 0; i < 10; i++ {
		channel <- i
		fmt.Println("send data ", i)
	}

	// close channel
	// close a channel to indicate that no more values will be sent
	close(channel)
}

func main() {

	// define new channel integer with length 2
	channel := make(chan int, 2)

	// run function goroutine
	go broadcast(channel)
	time.Sleep(2 * time.Second)

	// loop for read channel value
	for v := range channel {
		fmt.Println("read value", v, "from channel")
		time.Sleep(2 * time.Second)
	}

}
