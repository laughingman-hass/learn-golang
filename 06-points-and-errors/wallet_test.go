package wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		actual := wallet.Balance()
		expected := Bitcoin(10)

		if actual != expected {
			t.Errorf("actual %s, expected %s", actual, expected)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		actual := wallet.Balance()
		expected := Bitcoin(10)

		if actual != expected {
			t.Errorf("actual %s, expected %s", actual, expected)
		}
	})
}
