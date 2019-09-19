package main

import sh "./superhero"

func main() {
	thor := new(sh.Thor)
	superman := new(sh.Superman)

	sh.DoAttack(thor)
	sh.DoAttack(superman)
}
