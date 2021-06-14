package main

import (
	"fmt"
	"log"

	"github.com/suinautant/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(100)
	fmt.Println(account.Balance())
	err := account.Withdraw(120)
	if err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}
	fmt.Println(account.Balance())
}
