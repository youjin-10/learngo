package main

import (
	"fmt"

	"github.com/youjin-10/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("youjin")
	fmt.Println(account)
	fmt.Println(account.Owner())
	fmt.Println(account.Balance())
	account.Deposit(10)
	fmt.Println(account.Balance())
}


