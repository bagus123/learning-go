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

### Polymorphism

Polymorphism describes a pattern in object oriented programming in which classes have different functionality while sharing a common interface

#### Summary OOP in Go

- Encapsulation use Packages
- Inheritance use Composition
- Polymorphism use Interface
- Abstraction use Embedding

```

```
