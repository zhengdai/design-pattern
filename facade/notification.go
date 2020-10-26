package facade

import "fmt"

type notification struct {
}

func (n *notification) sendWalletCreditNotification() {
	fmt.Println("Sendinng wallet credit notification")
}

func (n *notification) sendWalletDebitNotification() {
	fmt.Println("Sendinng wallet debit notification")
}
