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
