package superhero

import "fmt"

type SuperHero interface {
	Attack()
}

func DoAttack(sh SuperHero) {
	sh.Attack()
}

type Superman struct{}

func (su Superman) Attack() {
	fmt.Println("Attach with laser")
}

type Thor struct{}

func (th Thor) Attack() {
	fmt.Println("Attach with Hammer")
}
