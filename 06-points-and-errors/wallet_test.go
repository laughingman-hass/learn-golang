package wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		actualError := wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, actualError)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		actualError := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, actualError, ErrInsufficientFunds)
	})
}

func assertBalance(t *testing.T, wallet Wallet, expected Bitcoin) {
	t.Helper()
	actual := wallet.Balance()

	if actual != expected {
		t.Errorf("actual %s, expected %s", actual, expected)
	}
}

func assertNoError(t *testing.T, actual error) {
	t.Helper()
	if actual != nil {
		t.Errorf("got an error but didn't want one")
	}
}

func assertError(t *testing.T, actualError error, expectedError error) {
	t.Helper()
	if actualError == nil {
		t.Errorf("wanted an error but didn't get one")
	}

	if actualError != expectedError {
		t.Errorf("actual %q, expected %q", actualError, expectedError)
	}
}
