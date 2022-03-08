package banking

import (
	"errors"
	"fmt"
)

type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("cat't withdraw")

func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

func (a *Account) SetBalance(amount int) {
	a.balance = amount
}

func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}

	a.balance -= amount
	return nil
}

func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) String() string {
	return fmt.Sprint("\n", a.owner, "'s account.\nHas:", a.balance, "\n")
}

func (a Account) Owner() string {
	return a.owner
}

func (a *Account) Balance() int {
	return a.balance
}
