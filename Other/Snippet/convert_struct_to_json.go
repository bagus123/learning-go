package main

import (
	"encoding/json"
	"fmt"
)

// Employee ...
type Employee struct {
	Name   string `json:"name"`
	Salary int    `json:"salary"`
}

func main() {
	employee := &Employee{Name: "John Doe", Salary: 10000000}
	json, err := json.Marshal(employee)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(json))
}
