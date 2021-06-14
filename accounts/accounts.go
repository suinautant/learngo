package accounts

import (
	"errors"
	"fmt"
)

// var errNoMoney = errors.New("Can't withdraw")

// Account struct - nico
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Acccount
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	fmt.Println("Gonna deposit : ", amount)
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("Can't withdraw")
		// return errors.New("Can't withdraw you are poor")
		// return errNoMoney
	} else {
		a.balance -= amount
		return nil
	}
}

// ChangeOwner of your account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	// return fmt.Sprintf(a.Owner(), "'s account.\nHas: ", a.Balance())
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
