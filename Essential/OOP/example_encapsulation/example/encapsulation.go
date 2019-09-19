package example

import "fmt"

type Encapsulation struct{}

// pubic : first letter must Uppercase
func (e *Encapsulation) Expose() {
	fmt.Println("this public function")
}

// private : first letter must lowercase
func (e *Encapsulation) hide() {
	fmt.Println("this private function")
}
