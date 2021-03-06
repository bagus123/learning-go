## Map

Map is Collection Object has key and value

#### Example 1 :

file example_map.go

```go
package main

import "fmt"

func main() {

    mapFruit := make(map[string]int)
    mapFruit["banana"] = 1
    mapFruit["orange"] = 2
    mapFruit["apple"] = 3

    // loop by for range type 1
    for key := range mapFruit {
    	fmt.Println("key", key, "value", mapFruit[key])
    }

    // loop by for range type 2
    for key, value := range mapFruit {
    	fmt.Println("key", key, "value", value)
    }

    // test if entry is present in the map or not
    value, exist := mapFruit["apple"]
    if exist {
    	fmt.Println("key is exist with value ", value)
    } else {
    	fmt.Println("key is not exist")
    }

    //delete an entry
    delete(mapFruit, "apple")
    valueAfter, exist := mapFruit["apple"]
    if exist {
    	fmt.Println("key is exist with value ", valueAfter)
    } else {
    	fmt.Println("key is not exist")
    }

}
```

#### Example 2:

file example_map_with_struct.go

```go
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
```
