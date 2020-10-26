package test

import (
	"design-pattern/facade"
	"fmt"
	"log"
	"testing"
)

func TestFacade(t *testing.T) {
	walletFacade := facade.NewWalletFacade("abc", 1234)
	fmt.Println()
	err := walletFacade.AddMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.DeductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
