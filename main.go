package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")

	const name string = "yj" // constant
	// var companyName string = "awesome studio" // variable
	// companyName := "awesome studio" // variable
	// companyName = "not awesome studio" // reassign variable

	fmt.Println(multiply(2, 3))

	length, upperCaseName := lenAndUpper("yj")
	fmt.Println(length, upperCaseName)

	shoutOut("yj", "123", "7890")

	totalNum := addAll(1, 2, 3, 4, 5)
	fmt.Println(totalNum)

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

func addAll(numbers ...int) int {
	total := 0
	// for i:=0; i<len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }
	for _, number := range numbers {
		total += number
	}
	return total
}

// a function that decides the parameter age is able to drink alcohol or not
func canIDrink(age int) bool {
	// create a variable inside if statement
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

