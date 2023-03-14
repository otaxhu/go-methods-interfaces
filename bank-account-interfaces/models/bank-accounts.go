package models

type BankAccount struct {
	Id       string `json:"id" gorm:"primaryKey;not null;type:varchar(255)"`
	UserId   string `json:"user_id" gorm:"not null;type:varchar(255)"`
	Balance  int    `json:"balance"`
	Currency string `json:"currency" gorm:"type:varchar(255)"`
}

func (b *BankAccount) Deposit(amount int) error {
	if amount <= 0 {
		return errInvalidAmount
	}
	b.Balance += amount
	return nil
}

func (b *BankAccount) Withdraw(amount int) error {
	if b.Balance-amount <= 0 {
		return errInsufficientFunds
	}
	b.Balance -= amount
	return nil
}

func (b *BankAccount) GetBalance() int {
	return b.Balance
}

func (b *BankAccount) GetCurrency() string {
	return b.Currency
}

func NewBolivaresAccount() *BankAccount {
	return &BankAccount{Currency: "Bolivares"}
}

func NewBitcoinAccount() *BankAccount {
	return &BankAccount{Currency: "Bitcoin"}
}

func NewUSDaccount() *BankAccount {
	return &BankAccount{Currency: "Dolares"}
}
