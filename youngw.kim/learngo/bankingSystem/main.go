package main

import (
	"fmt"
	"github.com/youngw.kim/learngo/bankingSystem/banking"
)

func main() {
	//account := banking.Account{Owner: "young", Balance: 10000}
	account := banking.NewAccount("young")
	fmt.Println(account)
	account.Deposit(40)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
	account.ChangeOwner("kyoung")
	fmt.Println(account.Owner())
	fmt.Println(account)

}
