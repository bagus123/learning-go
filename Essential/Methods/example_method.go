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
