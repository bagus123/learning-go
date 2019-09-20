package main

import oop "./oop"

func main() {

	superman := new(oop.Superman)
	thor := new(oop.Thor)

	s := oop.AbstractSuperHero{SuperHero: superman}
	t := oop.AbstractSuperHero{SuperHero: thor}

	// call abstract function MakeAttack
	s.MakeAttack()
	t.MakeAttack()

}
