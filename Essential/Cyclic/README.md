## Cyclic

example how to produce `import cycle not allowed`


```go

// file example_cyclic.go
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

// file animal/animal.go
package animal

import "log"

type Animal struct {
	Name string
}

func (a *Animal) Legs(num int) {

	if num == 2 {
		log.Println("walking with two legs")
	} else if num == 4 {
		log.Println("walking with four legs")
	}

}

// file people/people.go

package people

import (
	"log"

	p1 "../animal"
	p2 "../food"
)

type People struct {
	Name string
}

func (p *People) HavePet(animal *p1.Animal) {
	log.Println("have pet ", animal.Name)
	food := &p2.Food{Name: "Meat"}
	food.Taste(animal)
}

// file food/food.go

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



``
