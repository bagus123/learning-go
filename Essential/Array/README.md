## Array

Go programming language provides a data structure called the array, which can store a fixed-size sequential collection of elements of the same type. An array is used to store a collection of data, but it is often more useful to think of an array as a collection of variables of the same type

#### Example 1 :

file example_array.go

```go
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
```

### Slice

A slice is formed by specifying two indices, a low and high bound, separated by a colon:

```go
a[low : high]
```

This selects a half-open range which includes the first element, but excludes the last one.

The following expression creates a slice which includes elements 1 through 3 of a:

```go
a[1:4]
```

#### Example 2 :

file example_array_slice.go

```go
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
```

### Append

The built-in append function appends elements to the end of a slice:

- if there is enough capacity, the underlying array is reused;
- if not, a new underlying array is allocated and the data is copied over.

#### Example 3 :

file example_array_append.go

```go
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
```
