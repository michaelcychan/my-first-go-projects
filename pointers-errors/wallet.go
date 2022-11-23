package wallet

import (
	"errors"
	"fmt"
)

type Stringer interface {
	String() string
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (wal *Wallet) Deposit(amount Bitcoin) {
	(*wal).balance += amount
}

var ErrInsufficientFund = errors.New("cannot withdraw, insufficient fund")

func (wal *Wallet) Withdraw(amount Bitcoin) error {
	if wal.balance >= amount {
		wal.balance -= amount
		return nil
	}
	return ErrInsufficientFund
}

func (wal *Wallet) Balance() Bitcoin {
	return wal.balance
}
