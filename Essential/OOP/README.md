## OOP

### "Is Go object oriented?"

Yes, Go support for Object Oriented Programming
in OOP programing like Java or C++, they have class for create instance of Object, but in Go have 'struct'

### Encapsulation

Encapsulation is the mechanism of hiding of data implementation by restricting access to public methods

special notation in Go:

- function with first Letter Uppercase is public
- function with first letter lowercase is private

Example 1:

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

Example 1:

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

func (su Superman) Attack() {
	fmt.Println("Attach with laser")
}

type Thor struct{}

func (th Thor) Attack() {
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

	sh.DoAttack(thor)
	sh.DoAttack(superman)
}
```

#### Summary OOP in Go

- Encapsulation use Packages
- Inheritance use Composition
- Polymorphism use Interface
- Abstraction use Embedding

```

```
