package pointererrors

import (
	"errors"
	"fmt"
)

var InsufficientFundsException = errors.New("Insufficient funds")

type Stringer interface {
	String() string
}

type Bitcoin float64

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.4f BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return InsufficientFundsException
	}
	w.balance -= amount
	return nil
}
