package wallet

import (
	"testing"
)

/*
func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

}
*/

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	errorFunc := func(t *testing.T, got int, want int) {
		if got != want {
			t.Errorf("got %d wanted %d", got, want)
		}
	}

	t.Run("deposit into wallet", func(t *testing.T) {

		wallet.Deposit(10)
		got := wallet.Balance()
		//fmt.Printf("address in test is %p \n", &wallet.balance)
		want := 10

		errorFunc(t, got, want)
	})

}
