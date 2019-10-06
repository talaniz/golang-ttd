package pointers

import (
	"errors"
	"fmt"
)

// Bitcoin is the rep of a bitcoin
type Bitcoin int

// Wallet represents an account
type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Deposit adds funds to a wallet balance
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// ErrInsufficientFunds specifies the error for overdraft transactions
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw subtracts the amount from the balance
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

// Balance reports the balance of a wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
