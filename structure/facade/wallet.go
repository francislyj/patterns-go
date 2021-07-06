package facade

import "fmt"

type wallet struct {
	balance int
}

func newWallet() *wallet {
	return &wallet{
		balance: 0,
	}
}

func (w *wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Printf("credit balance %d \n", amount)
	return
}

func (w *wallet) debitAmount(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("balance is not sufficient")
	}
	fmt.Printf("debit wallet amount %d \n", amount)
	w.balance = w.balance - amount
	return nil

}