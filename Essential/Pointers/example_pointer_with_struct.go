package main

import "fmt"

type Car struct {
	name  string
	brand string
	model string
}

func main() {

	myCar := Car{
		"All New Vios",
		"Honda",
		"Vios",
	}

	p := &myCar

	fmt.Println("value myCar.brand is ", myCar.brand) // value myCar.brand is  Honda
	fmt.Println("value p.brand is ", p.brand)         // value p.brand is  Honda

	//To access the field brand of a struct when we have the struct pointer p we could write (*p).brand.
	//However, that notation is cumbersome,
	//so the language permits us instead to write just p.brand,
	//without the explicit dereference.

	// what if sign '&' not included
	// this is not assign address, but copy my car to copyOfCar
	copyOfCar := myCar
	copyOfCar.brand = "Toyota"

	fmt.Println("value myCar.brand is ", myCar.brand) // value myCar.brand is  Honda
	fmt.Println("value p.brand is ", copyOfCar.brand) // value p.brand is  Toyota
}
