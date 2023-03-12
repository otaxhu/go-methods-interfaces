package models

type bolivaresAccount struct {
	balance  int
	currency string
}

func NewBolivaresAccount() *bolivaresAccount {
	return &bolivaresAccount{currency: "Bolivares"}
}

func (b *bolivaresAccount) Deposit(amount int) error {
	if amount <= 0 {
		return errInvalidAmount
	}
	b.balance += amount
	return nil
}

func (b *bolivaresAccount) Withdraw(amount int) error {
	if b.balance-amount < 0 {
		return errInsufficientFunds
	}
	b.balance -= amount
	return nil
}

func (b *bolivaresAccount) GetBalance() int {
	return b.balance
}

func (b *bolivaresAccount) GetCurrency() string {
	return b.currency
}
