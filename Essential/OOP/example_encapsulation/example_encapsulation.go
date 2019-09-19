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
