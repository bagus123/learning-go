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
