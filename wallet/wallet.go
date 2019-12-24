package wallet

import "fmt"

type Wallet struct {
	balance int
}

func (w Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w Wallet) Balance() int {
	fmt.Printf("Address of the balance in deposit is %v\n", &w.balance)
	return w.balance
}
