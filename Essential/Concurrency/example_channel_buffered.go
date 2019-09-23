package main

import (
	"fmt"
)

func main() {

	// define new channel string with length 2
	channel := make(chan string, 2)

	// send 2x
	channel <- "satu"
	channel <- "dua"

	// receive 2x
	fmt.Println(<-channel)
	fmt.Println(<-channel)

}
