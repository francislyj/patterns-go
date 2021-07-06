package facade

import "log"

func TestFacade() {
	walletFacade := newWalletFacade("abc", 11)

	err := walletFacade.addMoneyToWallet("abc", 11, 20)
	if err != nil {
		log.Fatalf("Error is %s \n", err.Error())
	}

	err = walletFacade.debitMoneyFromWallet("abc", 11, 99)
	if err != nil {
		log.Fatalf("Error is %s \n", err.Error())
	}
}
