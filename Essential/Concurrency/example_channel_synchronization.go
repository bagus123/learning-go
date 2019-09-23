package main

import (
	"fmt"
	"time"
)

func worker(done chan bool, str string) {
	fmt.Print("working...")
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(str)
	}
	fmt.Println("done")
	done <- true
}

func main() {

	done := make(chan bool, 1)
	go worker(done, "goroutine")

	// If you removed the <- done line from this program, 
	// the program would exit before the worker even started.
	<-done
}
