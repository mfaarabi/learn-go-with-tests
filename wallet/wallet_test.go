package wallet

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))
	expected := Bitcoin(10)

	actual := wallet.Balance()

	if expected != actual {
		t.Errorf("Was expecting %s but got %s instead", expected, actual)
	}
}
