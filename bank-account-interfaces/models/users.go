package models

import (
	"github.com/otaxhu/go-methods-interfaces/bank-account-interfaces/interfaces"
)

type User struct {
	name     string
	accounts []interfaces.IBankAccounts
}

func NewUser() *User {
	return &User{}
}

func (u *User) SetName(name string) *User {
	u.name = name
	return u
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) AddBankAccounts(account ...interfaces.IBankAccounts) *User {
	u.accounts = append(u.accounts, account...)
	return u
}

func (u *User) GetBankAccounts() []interfaces.IBankAccounts {
	return u.accounts
}

func (u *User) Deposit(amount int, currency string) error {
	for _, account := range u.accounts {
		if account.GetCurrency() == currency {
			if err := account.Deposit(amount); err != nil {
				return err
			}
			return nil
		}
	}
	return errCurrencyAccountNotFound
}

func (u *User) Withdraw(amount int, currency string) error {
	for _, account := range u.accounts {
		if account.GetCurrency() == currency {
			if err := account.Withdraw(amount); err != nil {
				return err
			}
			return nil
		}
	}
	return errCurrencyAccountNotFound
}
