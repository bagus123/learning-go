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
