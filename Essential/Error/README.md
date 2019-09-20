## Error

Go programming provides a pretty simple error handling framework with inbuilt error interface type of the following declaration −

```go
type error interface {
   Error() string
}
```

By convention, errors are the last return value and have type error, a built-in interface.

errors.New constructs a basic error value with the given error message

A nil value in the error position indicates that there was no error

#### Example 1 :

file example_error.go

```go
package main

import (
	"errors"
	"fmt"
)

func divided(a, b int) (int, error) {

	if b == 0 {
		return -1, errors.New("can't divided by zero")
	}

	c := a / b
	return c, nil
}

func main() {

	result1, err := divided(12, 3)

	if err != nil {
		fmt.Println("error ", err)
	} else {
		fmt.Println("result1 is ", result1) // result1 is  4
	}

	result2, err := divided(3, 0)

	if err != nil {
		fmt.Println("error ", err) //error can't divided by zero
	} else {
		fmt.Println("result2 is ", result2)
	}
}
```
