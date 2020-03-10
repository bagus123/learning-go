package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	data, err := ioutil.ReadFile("/Users/dev/GoProjects/learning-go/Essential/ReadFile/my_file.txt")
	if err != nil {
		fmt.Println("reading file error", err)
	}

	fmt.Println("contents of file", string(data))
}
