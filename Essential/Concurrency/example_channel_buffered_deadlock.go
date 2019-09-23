package main

import (
	"fmt"
)

func main() {

	// define new channel string with length 2
	channel := make(chan string, 2)
	channel <- "satu"
	channel <- "dua"

	// when channel try to send 3rd message will error/deadlock
	channel <- "tiga"

	// why error/ deadlock ?
	// because channel only reserve 2,
	// 2x send, 2x receive

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
