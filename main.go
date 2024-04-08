package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	// map
	personMap := map[string]string{"name": "Q", "age": "20"}
	fmt.Println(personMap)
	for key, index := range personMap {
		fmt.Println(key, index)
	}
}


