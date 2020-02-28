## Function

### What is a Function?

A function is a group of statements that together perform a task. Every Go program has at least one function, which is main(). You can divide your code into separate functions. How you divide your code among different functions is up to you, but logically, the division should be such that each function performs a specific task.

A function declaration tells the compiler about a function name, return type, and parameters. A function definition provides the actual body of the function.

The Go standard library provides numerous built-in functions that your program can call. For example, the function len() takes arguments of various types and returns the length of the type. If a string is passed to it, the function returns the length of the string in bytes. If an array is passed to it, the function returns the length of the array.

Functions are also known as method, sub-routine, or procedure.


#### Example 1 :

```go
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

```