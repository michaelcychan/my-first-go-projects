package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		expectedBalance := Bitcoin(10)

		assertBalance(t, wallet, expectedBalance)

	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw when insufficient fund", func(t *testing.T) {
		initialBalance := Bitcoin(20)
		wallet := Wallet{balance: initialBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFund)
		assertBalance(t, wallet, initialBalance)

	})
}

func assertBalance(t testing.TB, wall Wallet, expectedBalance Bitcoin) {
	t.Helper()
	if wall.Balance() != expectedBalance {
		t.Errorf("expected %s, but got %s", expectedBalance, wall.Balance())
	}
}

func assertError(t testing.TB, got error, expected error) {
	t.Helper()
	if got == nil {
		t.Fatal("Expected an error, but did not get one") // t.Fatal will stop the whole test
	}
	if got != expected {
		t.Errorf("Expected %q, but got %q", expected, got)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("Unexpected error occured")
	}
}
