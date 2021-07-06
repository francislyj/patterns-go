package facade

import "fmt"

type notification struct {
	
}


func (n *notification) sendWalletCreditNotification() {
	fmt.Println("send wallet credit notification")
}

func (n *notification) sendWalletDebitNotification() {
	fmt.Println("send wallet debit notification")
}