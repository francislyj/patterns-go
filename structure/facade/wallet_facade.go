package facade

import "fmt"

type walletFacade struct {
	account *account
	wallet *wallet
	securityCode *securityCode
	notification *notification
	ledger *ledger
}

func newWalletFacade(accountID string, code int) *walletFacade {
	fmt.Println("begin create wallet")
	walletFacade := &walletFacade{
		account: newAccount(accountID),
		securityCode: newSecurityCode(code),
		wallet: newWallet(),
		notification: &notification{},
		ledger: &ledger{},
	}
	fmt.Println("wallet has created")
	return walletFacade
}

func (w *walletFacade) addMoneyToWallet(accountID string, securityCode int, amount int) error {

	fmt.Println("starting add money to wallet")

	err := w.account.checkAccount(accountID)

	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode)

	if err != nil {
		return err
	}

	w.wallet.creditBalance(amount)
	w.notification.sendWalletDebitNotification()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}

func (w *walletFacade) debitMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("begin debit money from wallet")

	err := w.account.checkAccount(accountID)
	if err != nil {
		return nil
	}

	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}

	err = w.wallet.debitAmount(amount)
	if err != nil {
		return err
	}

	w.notification.sendWalletDebitNotification()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}