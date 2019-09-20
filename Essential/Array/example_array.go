package main

import "fmt"

func main() {

	// style 1: define array
	var numbers [6]int
	for i := 0; i < 6; i++ {
		numbers[i] = i
	}
	fmt.Println("numbers ", numbers)

	// style 2 : define array and init
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("primes", primes)
}
