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
