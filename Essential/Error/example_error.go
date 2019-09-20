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
