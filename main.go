package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	// with length: array
	numbers := [3]int{1, 2, 3}
	fmt.Println(numbers)
	// without length: slice
	nums := []int{1, 2, 3}
	nums = append(nums, 4)
	fmt.Println(nums)
	
}


