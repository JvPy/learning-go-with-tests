package pointers

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds") //this makes ErrInsufficientFunds global to the package

type Bitcoin int

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin //this is a private prop outside this package
}

//Deposit - makes a deposit to your account
func (w *Wallet) Deposit(amount Bitcoin) {
	// fmt.Printf("address of balance is Deposit is: %v\n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

//Balance - check your accounts balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
