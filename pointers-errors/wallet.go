package wallet

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (wal *Wallet) Deposit(money Bitcoin) {
	fmt.Printf("address of balance in function is %v \n", &wal.balance)
	(*wal).balance += money
}

func (wal *Wallet) Balance() Bitcoin {
	return wal.balance
}
