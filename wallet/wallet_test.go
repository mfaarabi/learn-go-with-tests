package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)
	expected := 10

	actual := wallet.Balance()

	assert.Equal(t, expected, actual)
}
