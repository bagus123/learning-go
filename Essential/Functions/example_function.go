package main

import "fmt"

// this function run first
func init() {

	fmt.Println("run first")
}

// this variadic function
func calculate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}

func Add(x int, y int) int {
	return x + y
}

// this main function
func main() {

	fmt.Println("run after init")
	fmt.Printf("result is %v \n", Add(1, 2))

	var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
	fmt.Println(msg)

}
