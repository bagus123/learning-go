package oop

import "fmt"

// SuperHero is interface
type SuperHero interface {
	Attack()
}

// AbstractSuperHero is abstract of SuperHero
type AbstractSuperHero struct {
	SuperHero
}

// MakeAttack is abstract function
func (a AbstractSuperHero) MakeAttack() {
	a.Attack()
}

// Superman is new struct
type Superman struct{}

// Attack is implement from interface SuperHero
func (s *Superman) Attack() {
	fmt.Println("Attack with laser")
}

// Thor is new struct
type Thor struct{}

// Attack is implement from interface SuperHero
func (t *Thor) Attack() {
	fmt.Println("Attack with Hammer")
}
