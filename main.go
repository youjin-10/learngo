package main

import (
	"fmt"

	"github.com/youjin-10/learngo/mydict"
)

func main() {
	myDictionary := mydict.Dictionary{"first": "First word"}

	word, err := myDictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
	
	err2 := myDictionary.Add("water", "something that you drink")
	fmt.Println(err2)

	

}


