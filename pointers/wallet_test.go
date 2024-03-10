package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	startingBalance := Bitcoin(100)
	movingMoney := Bitcoin(10)

	t.Run("deposit into wallet", func(t *testing.T) {
		wallet := Wallet{startingBalance}
		wallet.Deposit(movingMoney)

		balanceFunc(t, wallet, startingBalance+movingMoney)
	})
	//withdraw amount defined as starting + deposit
	t.Run("withdraw with sufficient funds", func(t *testing.T) {
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(movingMoney)
		assertNoError(t, err)
		balanceFunc(t, wallet, startingBalance-movingMoney)
	})

	t.Run("withdraw with insufficient funds", func(t *testing.T) {
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(startingBalance + movingMoney)

		throwError(t, err, ErrInsufficientFunds)
		balanceFunc(t, wallet, startingBalance)

	})
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func throwError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("Didnt catch the error")
	}

	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}

func balanceFunc(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s wanted %s", got, want)
	}
}
