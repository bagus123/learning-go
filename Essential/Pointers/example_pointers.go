package main

import "fmt"

func main() {

	// define new variable integer
	var a int
	a = 10

	// define new variable pointer integer
	var p *int

	// assign address a to pointer p
	p = &a

	fmt.Println("value a", a)
	fmt.Println("value p", *p)
	fmt.Println("address p", p)

	// shortcut define new variable with assign number
	i := 120

	// shortcut define new variable pointer with assign address of i
	k := &i

	fmt.Println("value i", i)
	fmt.Println("value k", *k)
	fmt.Println("address k", k)
}
