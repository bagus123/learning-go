package superhero

import "fmt"

// SuperHero is interface
type SuperHero interface {
	Attack()
}

// DoAttack : public function
func DoAttack(sh SuperHero) {
	sh.Attack()
}

// Superman struct
type Superman struct{}

// Attack : implement interface SuperHero
func (su Superman) Attack() {
	fmt.Println("Attach with laser")
}

// Thor struct
type Thor struct{}

// Attack : implement interface SuperHero
func (th Thor) Attack() {
	fmt.Println("Attach with Hammer")
}
