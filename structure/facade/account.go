package facade

import "fmt"

type account struct {
	name string
}

func newAccount(accountName string) *account {
	return &account{
		name: accountName,
	}
}

func (a *account) checkAccount(accountID string) error {
	if a.name != accountID {
		return fmt.Errorf("account name is error")
	}
	fmt.Println("account verified")
	return nil
}
