package interfaces

type IBankAccounts interface {
	Deposit(amount int) error
	Withdraw(amount int) error
	GetBalance() int
	GetCurrency() string
}
