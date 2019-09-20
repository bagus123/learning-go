package main

import "fmt"

func main() {

	// style 1: define array
	var numbers [6]int
	for i := 0; i < 6; i++ {
		numbers[i] = i
	}

	fmt.Println("numbers ", numbers)
	fmt.Println("slice numbers from index 2 to length", numbers[2:])
	fmt.Println("slice numbers from index 0 to 4", numbers[:5])
	fmt.Println("slice numbers from index 2 to 4", numbers[2:5])
}
