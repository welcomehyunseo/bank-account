package main

import (
	"fmt"

	"github.com/welcomehyunseo/bank-account/banking"
)

func main() {
	account := banking.Account{Owner: "hyunseo", Balance: 1000}
	fmt.Println(account)
}
