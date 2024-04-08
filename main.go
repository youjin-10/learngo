package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	a := 2
	b := &a
	*b = 123
	fmt.Println(a, *b)
}


