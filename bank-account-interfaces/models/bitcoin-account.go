package models

type bitcoinAccount struct {
	balance  int
	currency string
}

func NewBitcoinAccount() *bitcoinAccount {
	return &bitcoinAccount{currency: "Bitcoin"}
}

func (b *bitcoinAccount) Deposit(ammount int) error {
	if ammount <= 0 {
		return errInvalidAmount
	}
	b.balance += ammount
	return nil
}

func (b *bitcoinAccount) Withdraw(ammount int) error {
	if b.balance-ammount < 0 {
		return errInsufficientFunds
	}
	b.balance -= ammount
	return nil
}

func (b *bitcoinAccount) GetBalance() int {
	return b.balance
}

func (b *bitcoinAccount) GetCurrency() string {
	return b.currency
}
