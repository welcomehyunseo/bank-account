package main

import (
	"fmt"
	"log"

	"github.com/welcomehyunseo/bank-account/banking"
)

func main() {
	account := banking.NewAccount("hyunseo")
	account.SetBalance(10)
	fmt.Println(account.String())

	// error handling in golang
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(account.String())
}
