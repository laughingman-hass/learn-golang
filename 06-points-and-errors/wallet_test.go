package wallet

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, expected Bitcoin) {
		t.Helper()
		actual := wallet.Balance()

		if actual != expected {
			t.Errorf("actual %s, expected %s", actual, expected)
		}
	}

	assertError := func(t *testing.T, actualError error, expectedError error) {
		t.Helper()
		if actualError == nil {
			t.Errorf("wanted an error but didn't get one")
		}

		if actualError != expectedError {
			t.Errorf("actual %q, expected %q", actualError, expectedError)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		actualError := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, actualError, ErrInsufficientFunds)
	})
}
