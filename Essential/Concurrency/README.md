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
