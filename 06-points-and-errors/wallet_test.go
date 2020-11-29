package wallet

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	actual := wallet.Balance()
	expected := 10

	if actual != expected {
		t.Errorf("actual %d, expected %d", actual, expected)
	}
}
