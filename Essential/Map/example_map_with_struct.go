package main

import "fmt"

type Employee struct {
	Name     string
	Position string
}

func main() {

	// define map type 1
	mapEmployee := make(map[string]Employee)
	mapEmployee["001"] = Employee{
		Name:     "John Doe",
		Position: "Manager",
	}

	mapEmployee["002"] = Employee{
		Name:     "John Paul",
		Position: "Supervisor",
	}

	for _, value := range mapEmployee {
		fmt.Println("Employee ", value)
	}

	// define map type 2
	mapEmployee2 := map[string]Employee{
		"001": Employee{
			Name:     "John Doe",
			Position: "Manager",
		},
		"002": Employee{
			Name:     "John Paul",
			Position: "Supervisor",
		},
	}

	for _, value := range mapEmployee2 {
		fmt.Println("Employee ", value)
	}

	// define map type 3
	mapEmployee3 := map[string]Employee{}

	mapEmployee3["001"] = Employee{
		Name:     "John Doe",
		Position: "Manager",
	}

	mapEmployee3["002"] = Employee{
		Name:     "John Paul",
		Position: "Supervisor",
	}

	for _, value := range mapEmployee3 {
		fmt.Println("Employee ", value)
	}

}
