package food

import (
	"log"

	p1 "../animal"
)

type Food struct {
	Name string
}

func (f *Food) Taste(animal *p1.Animal) {
	if f.Name == "Meat" && animal.Name == "Lion" {
		log.Println("Yummy")

		// code detected as cyclic
		//people :=&p2.People(Name:"John")
		//people.HavePet(animal)

	}
}
