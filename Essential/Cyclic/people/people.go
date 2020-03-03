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
