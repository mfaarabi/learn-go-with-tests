package wallet

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, expected, actual Bitcoin) {
		t.Helper()
		if expected != actual {
			t.Errorf("Was expecting %s but got %s instead", expected, actual)
		}
	}

	assertError := func(t *testing.T, expected string, actual error) {
		t.Helper()
		if actual == nil {
			t.Fatal("Was expecting an error but didn't get one")
		}
		if actual.Error() != expected {
			t.Errorf("Was expecting %q but got %q instead", expected, actual)
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

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		expected := Bitcoin(20)

		err := wallet.Withdraw(Bitcoin(100))
		actual := wallet.Balance()

		assertBalance(t, expected, actual)
		assertError(t, "Amount to withdraw is exceeding balance", err)
	})
}
