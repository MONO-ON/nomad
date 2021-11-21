package main

import (
	"fmt"

	"github.com/MONO-ON/nomad/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("mono")
	account.Deposit(10)
	fmt.Println(account)
}
