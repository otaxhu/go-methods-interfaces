package models

type usdAccount struct {
	balance  int
	currency string
}

func NewUSDaccount() *usdAccount {
	return &usdAccount{currency: "Dolares"}
}

func (u *usdAccount) Deposit(amount int) error {
	if amount <= 0 {
		return errInvalidAmount
	}
	u.balance += amount
	return nil
}

func (u *usdAccount) Withdraw(amount int) error {
	if u.balance-amount < 0 {
		return errInsufficientFunds
	}
	u.balance -= amount
	return nil
}

func (u *usdAccount) GetBalance() int {
	return u.balance
}

func (u *usdAccount) GetCurrency() string {
	return u.currency
}
