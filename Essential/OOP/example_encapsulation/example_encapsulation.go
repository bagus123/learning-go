package main

// import with alias c
import c "./example"

func main() {
	// create instance Encapsulation
	e := c.Encapsulation{}

	// run public function Expose
	e.Expose()

	// function hide not exposed, this function as private
	// e.hide() // if this uncomment will throw error : .\example_encapsulation.go:14:3: e.hide undefined (cannot refer to unexported field or method example.(*Encapsulation).hide)

}
