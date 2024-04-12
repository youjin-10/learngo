package main

import (
	"fmt"

	"github.com/youjin-10/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("youjin")
	fmt.Println(account.Owner())
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(25)
	if err != nil {
		// log.Fatalln(err)
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}


