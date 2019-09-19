package main

import "fmt"

type SuperHeroes struct {
	name    string
	ability string
	power   int
}

func main() {

	superman := &SuperHeroes{
		"superman",
		"eye of laser",
		100,
	}

	spiderman := &SuperHeroes{
		"spiderman",
		"power of webs",
		100,
	}

	fmt.Println("superman has ability ", superman.ability)
	fmt.Println("spiderman has ability ", spiderman.ability)

}
