## OOP

### "Is Go object oriented?"

Yes, Go support for Object Oriented Programming
in OOP programing like Java or C++, they have class for create instance of Object, but in Go has 'struct'

### Encapsulation

Encapsulation is the mechanism of hiding of data implementation by restricting access to public methods

special notation in Go:

- function with first Letter Uppercase is public
- function with first letter lowercase is private

#### Example 1:

file encapsulation.go

```go
package example


import "fmt"

type Encapsulation struct{}

// pubic function : first letter must Uppercase
func (e *Encapsulation) Expose() {
	fmt.Println("this public function")
}

// private function : first letter must lowercase
func (e *Encapsulation) hide() {
	fmt.Println("this private function")
}
```

file example_encapsulation.go

```go
package main

// import with alias oop
import oop "./oop"

func main() {
	// create instance Encapsulation
	e := oop.Encapsulation{}

	// run public function Expose
	e.Expose()

	// function hide not exposed, this function as private
	// e.hide() // if this uncomment will throw error : .\example_encapsulation.go:14:3: e.hide undefined (cannot refer to unexported field or method example.(*Encapsulation).hide)

}
```

### Polymorphism

Polymorphism describes a pattern in object oriented programming in which classes have different functionality while sharing a common interface

#### Example 1:

file polymorphism.go

```go
package oop

import "fmt"

type SuperHero interface {
	Attack()
}

func DoAttack(sh SuperHero) {
	sh.Attack()
}

type Superman struct{}

func (su *Superman) Attack() {
	fmt.Println("Attach with laser")
}

type Thor struct{}

func (th *Thor) Attack() {
	fmt.Println("Attach with Hammer")
}
```

file example_polymorphism.go

```go
package main

import sh "./superhero"

func main() {
	thor := new(sh.Thor)
	superman := new(sh.Superman)

	sh.DoAttack(thor)     // Attach with Hammer
	sh.DoAttack(superman) // Attach with laser
}
```

### Composition

In Go, inheritance is not possible. Instead, we build our structs with composable and reusable elements through embedding.

Go allows us to embed types within interfaces or structs. Through embedding, we are able to forward the methods included from the inner type, to the outer type.

#### Example 1:

file composition.go

```go
package oop

import (
	"fmt"
)

// Vehicle is primary struct
// don't forget use Uppercase first letter for access public
type Vehicle struct {
	Name  string
	Brand string
}

// Car composed Vehicle with NumberOfDoor, NumberOfPassager
type Car struct {
	Vehicle
	NumberOfDoor     int
	NumberOfPassager int
}

// MotorCycle composed Vehicle with MotorCycleType
type MotorCycle struct {
	Vehicle
	MotorCycleType string
}

// Show ...
func (v *Vehicle) Show() {
	fmt.Println("this Brand of vehicle is ", v.Brand)
}

```

file example_composition.go

```go
package main

import (
	v "./oop"
)

func main() {
	car := v.Car{
		Vehicle: v.Vehicle{
			Name:  "Toyota Vios",
			Brand: "Toyota",
		},
		NumberOfDoor:     4,
		NumberOfPassager: 7,
	}

	motorCycle := v.MotorCycle{
		Vehicle: v.Vehicle{
			Name:  "Honda CBR 250",
			Brand: "Honda",
		},
		MotorCycleType: "Sport",
	}

	car.Show()
	motorCycle.Show()
}
```

### Abstraction

Abstraction means working with something we know how to use without knowing how it works internally.

Similar to embedding structs within a struct, we can also embed interfaces within structs. Remember that any type that satisfied an interface also adopts that interface type.

#### Example 1:

file abstraction.go

```go
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
```

file example_abstraction.go

```go
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

```

#### Summary OOP in Go

- Encapsulation use Packages
- Inheritance use Composition
- Polymorphism use Interface
- Abstraction use Embedding

```

```
