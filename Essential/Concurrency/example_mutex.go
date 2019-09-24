package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {

	// lock mutex
	m.Lock()
	x = x + 1

	// unlock mutex
	m.Unlock()

	// release wait
	wg.Done()
}
func main() {

	// define waitGroup
	var w sync.WaitGroup

	// define mutex
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		// add numbers goroutine will be wait (block)
		w.Add(1)
		// run goroutine
		go increment(&w, &m)
	}

	// wait block until done
	w.Wait()

	// print value x
	fmt.Println("final value of x", x)
}
