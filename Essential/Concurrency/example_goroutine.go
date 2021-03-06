package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {

	// set max core CPU for execute
	runtime.GOMAXPROCS(2)

	// run multi thread using goroutine (asynchronous)
	// go functionName
	go say("hello goroutine")

	// call direct function
	say("hello direct function")

	// goroutine can write like this, but run not asynchronous
	//
	// go func(s string) {
	// 	for i := 0; i < 10; i++ {
	// 		time.Sleep(100 * time.Millisecond)
	// 		fmt.Println(s)
	// 	}
	// }("hello goroutine inner function")
	// fmt.Scanln()
	// fmt.Println("done")
}
