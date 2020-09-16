package banking

import (
	"errors"
	"fmt"
)

// account struct
type account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

// create new account function
func NewAccount(owner string) *account {
	nAccount := account{owner: owner, balance: 0}
	return &nAccount
}

// Deposit to your account
func (a *account) Deposit(amount int) {
	a.balance += amount
}

// check balance of your account
func (a account) Balance() int {
	return a.balance
}

func (a *account) Withdraw(amount int) error {
	if a.balance < amount {
		//return errors.New("Cannot Withdraw")
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

func (a *account) ChangeOwner(owner string) {
	a.owner = owner
}

func (a account) Owner() string {
	return a.owner
}

func (a account) String() string {
	return fmt.Sprint("-----------\nOwner name : ", a.owner, "\nRemain : ", a.balance)
}
