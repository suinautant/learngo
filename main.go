package main

import (
	"fmt"

	"github.com/suinautant/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)
}
