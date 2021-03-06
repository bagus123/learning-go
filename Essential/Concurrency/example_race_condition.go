package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1

	// release block
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {

		// add numbers goroutine will be wait (block)
		w.Add(1)

		// run goroutine
		go increment(&w)
	}

	// wait block until done
	w.Wait()

	// print value x
	fmt.Println("final value of x", x)
	// output
	// final value of x 1000
	// final value of x 992
	// final value of x 92
}
