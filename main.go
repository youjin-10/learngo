package main

import (
	"fmt"
)

type person struct {
	name      string
	age       int
	isStudent bool
	favFruit  []string
}

func main() {
	fmt.Println("Hello, World!")
	// map
	personMap := map[string]string{"name": "Q", "age": "20"}
	fmt.Println(personMap)
	for key, index := range personMap {
		fmt.Println(key, index)
	}

	/*
	{
		"name": "Q",
		"age": 20,
		"isStudent": true,
		"favFruit": ["apple", "banana"],
	}
	*/
	// struct
	p := person{
		name:      "Q",
		age:       20,
		isStudent: true,
		favFruit:  []string{"apple", "banana"},
	}
	fmt.Println(p)
	

}


