package wallet

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, expected, actual Bitcoin) {
		t.Helper()
		if expected != actual {
			t.Errorf("Was expecting %s but got %s instead", expected, actual)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		expected := Bitcoin(10)

		actual := wallet.Balance()

		assertBalance(t, expected, actual)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		expected := Bitcoin(10)

		actual := wallet.Balance()

		assertBalance(t, expected, actual)
	})
}
