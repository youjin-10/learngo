package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")

	const name string = "yj" // constant
	// var companyName string = "awesome studio" // variable
	companyName := "awesome studio" // variable
	companyName = "not awesome studio" // reassign variable

	fmt.Println(multiply(2, 3))

	length, upperCaseName := lenAndUpper("yj")
	fmt.Println(length, upperCaseName)

	shoutOut("yj", "123", "7890")
}

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	defer fmt.Println("finished!") // defer: after function is done, run this line
	return len(name), strings.ToUpper(name)
}

func shoutOut(names ...string) {
	fmt.Println(names)
}