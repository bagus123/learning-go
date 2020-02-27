package main

import (
	"fmt"
	"log"
)

type SuperHeroes struct {
	name    string
	ability string
	power   int
}

func main() {

	// var superman is pointer, assign address by ampersand (&)
	superman := &SuperHeroes{
		"superman",
		"eye of laser",
		100,
	}

	// var spiderman is pointer, assign address by ampersand (&)
	spiderman := &SuperHeroes{
		"spiderman",
		"power of webs",
		100,
	}

	fmt.Println("superman has ability ", superman.ability)
	fmt.Println("spiderman has ability ", spiderman.ability)

	x := 1
	y := &x

	log.Printf("value y is %v (address) ", y)
	log.Printf("value *y is %v", *y)

}
