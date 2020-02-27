## Methods

A method is a function with a special receiver argument.

The receiver appears in its own argument list between the func keyword and the method name.

In this example, the Abs method has a receiver of type Vertex named v.


```go

package main

import (
	"fmt"
)

type MathOps struct {
	X, Y float64
}

// this function is owner of MathOps struct
func (v MathOps) Add() float64 {
	return v.X + v.Y
}

// this function is owner of MathOps struct
func main() {
	v := MathOps{3, 4}

	fmt.Println(v.X)     // output : 3
	fmt.Println(v.Y)     // output : 4
	fmt.Println(v.Add()) // output : 7
}


```