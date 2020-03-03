package main

import (
	"log"

	p1 "./animal"
	p2 "./people"
)

func main() {

	animal := &p1.Animal{Name: "Lion"}
	people := &p2.People{Name: "John Doe"}
	people.HavePet(animal)

	log.Println(animal)
	log.Println(people)

}
