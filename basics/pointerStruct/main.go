package main

import "fmt"

type Account struct {
	owner string
	value int
}

func (a Account) simulateLoan(loan int) int {
	a.value += loan
	return a.value
}

func (a *Account) credit(value int) int {
	a.value += value
	return a.value
}

func newAccount() *Account {
	return &Account{value: 0}
}

func main() {
	wendAccount := Account{owner: "Wend", value: 100}

	fmt.Printf("Loan simulation of 50 for account is: %v\n", wendAccount.simulateLoan(50))
	fmt.Printf("Account after simulation: %v\n", wendAccount.value)
	fmt.Printf("Credit of 50 for account is: %v\n", wendAccount.credit(50))
	fmt.Printf("Account after credit: %v\n", wendAccount.value)

	account := newAccount()

	account.credit(50)

	println(account.value)
}
