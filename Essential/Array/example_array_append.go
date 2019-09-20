package main

import "fmt"

func main() {

	// if dynamic array don't fill size of array
	numbers := []int{1, 2, 3, 4, 5, 6}

	fmt.Println("numbers", numbers) // numbers [1 2 3 4 5 6]

	for i := 0; i < 4; i++ {
		newSize := len(numbers)
		numbers = append(numbers, newSize+1) // numbers after append [1 2 3 4 5 6 7 8 9 10]

	fmt.Println("numbers after append", numbers)
}
