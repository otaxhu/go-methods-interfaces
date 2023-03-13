package main

import (
	"fmt"

	"github.com/otaxhu/go-methods-interfaces/bank-account-interfaces/interfaces"
	"github.com/otaxhu/go-methods-interfaces/bank-account-interfaces/models"
)

func main() {
	//router.InitializeRouter()
	bolivares := models.NewBolivaresAccount()
	dolares := models.NewUSDaccount()
	user1 := models.NewUser().SetName("Oscar Pernia").AddBankAccounts(bolivares, dolares)
	printUserWithBalances(user1)
	bolivares.Deposit(100)
	printUserWithBalances(user1)
}

func printBalance(account interfaces.IBankAccounts) {
	fmt.Printf("Usted tiene %d en su cuenta de %s\n", account.GetBalance(), account.GetCurrency())
}

func printUserWithBalances(user *models.User) {
	for _, account := range user.GetBankAccounts() {
		fmt.Printf("este usuario %s tiene %d en su cuenta de %s\n", user.GetName(), account.GetBalance(), account.GetCurrency())
	}
}
