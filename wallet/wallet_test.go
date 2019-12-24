package wallet

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)
	expected := 10

	actual := wallet.Balance()
	fmt.Printf("Address of the balance in test is %v\n", &wallet.balance)

	assert.Equal(t, expected, actual)
}
