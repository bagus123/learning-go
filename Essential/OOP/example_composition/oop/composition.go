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
