package main

import sh "./superhero"

func main() {
	thor := new(sh.Thor)
	superman := new(sh.Superman)

	sh.DoAttack(thor)     // Attach with Hammer
	sh.DoAttack(superman) // Attach with laser
}
