package wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	gotBalance := wallet.Balance()

	fmt.Printf("address of balance in test is %v \n", &wallet.balance)

	expectedBalance := Bitcoin(10)

	if gotBalance != expectedBalance {
		t.Errorf("expected %d, but got %d", expectedBalance, gotBalance)
	}
}
