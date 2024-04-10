package main

import (
	"fmt"

	"github.com/youjin-10/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("youjin")
	fmt.Println(account)
}


