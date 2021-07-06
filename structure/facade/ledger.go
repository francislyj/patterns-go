package facade

import "fmt"

type ledger struct {
	
}

func (l *ledger) makeEntry(accountID, txnType string, amount int) {
	fmt.Printf("make ledger entry for accountID %s with txnType %s for amount %d \n", accountID, txnType, amount)
	return
}
