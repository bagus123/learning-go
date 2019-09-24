## Concurrency

### Goroutines

Go has goroutine for manage multi thread like Java, goroutine define very simple

Define goroutine

```go
go f(x, y, z)
```

#### Example 1:

file example_goroutine.go

```go
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {

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
```

### Channel

Channels are a typed conduit through which you can send and receive values with the channel operator, <-.

```go
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
```

Define Channel

```go
ch := make(chan int)
```

#### Example 1:

file example_channel.go

```go
package main

func sum(numbers []int, channel chan int) {
	result := 0
	for _, value := range numbers {
		result += value
	}

	// send to channel
	channel <- result
}

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// define new channel type integer
	channel := make(chan int)

	// run function using goroutine (Thread)
	go sum(numbers, channel)

	// receive from channel
	result := <-channel

	println("The Result is ", result)
}
```

### Buffered Channels

it's channel with size of channel, sends to a buffered channel are blocked only when the buffer is full. Similarly receives from a buffered channel are blocked only when the buffer is empty.

define buffered channel

```go
ch := make(chan int, 100)
```

#### Example 1:

file example_channel_buffered.go

```go
package main

import (
	"fmt"
)

func main() {

    // define new channel string with length 2
    channel := make(chan string, 2)

    // send 2x
    channel <- "satu"

    // if comment this code will throw error in line [A]
    channel <- "dua"

    // receive 2x
    fmt.Println(<-channel)

    // [A]
    fmt.Println(<-channel)

}
```

#### Example 2:

file example_channel_buffered_loop.go

```go
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
```

#### Example 3:

file example_channel_buffered_deadlock.go

```go
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
```

#### Example 4:

file example_channel_synchronization.go

```go
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
```

### Select

Go’s select lets you wait on multiple channel operations. Combining goroutines and channels with select is a powerful feature of Go.

select is just like switch without any input argument but it only used for channel operations. The select statement is used to perform an operation on only one channel out of many, conditionally selected by case block.

#### Example 1:

file example_select.go

```go
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
```

#### default case

Like switch statement, select statement also has default case. A default case is non-blocking. But that’s not all, default case makes select statement always non-blocking. That means, send and receive operation on any channel (buffered or unbuffered) is always non-blocking.

#### Example 2:

file example_select_non_blocking.go

```go
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
```

### Mutex

A Mutex (mutual exclusion) is used to provide a locking mechanism to ensure that only one Goroutine is running the critical section of code at any point of time to prevent race condition from happening.

#### Race Condition

#### Example 1:

file example_race_condition.go

```go
package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w)
	}
	w.Wait()
	fmt.Println("final value of x", x)
	// output
	// final value of x 1000
	// final value of x 992
	// final value of x 92
}
```

the output will be different for each time because of race condition. Some of the outputs which I encountered are final value of x 941, final value of x 928, final value of x 922 and so on

> **How to detect race condition using command,
> use this command**

```go
go run -race filename.go
```

for example :

```go
D:\learning-go\Essential\Concurrency>go run -race example_race_condition.go
==================
WARNING: DATA RACE
Read at 0x000000607318 by goroutine 7:
  main.increment()
      D:/learning-go/Essential/Concurrency/example_race_condition.go:11 +0x41

Previous write at 0x000000607318 by goroutine 6:
  main.increment()
      D:/learning-go/Essential/Concurrency/example_race_condition.go:11 +0x5d

Goroutine 7 (running) created at:
  main.main()
      D:/learning-go/Essential/Concurrency/example_race_condition.go:18 +0xb2

Goroutine 6 (finished) created at:
  main.main()
      D:/learning-go/Essential/Concurrency/example_race_condition.go:18 +0xb2
==================
final value of x 999
Found 1 data race(s)  // race detected
exit status 66
```
