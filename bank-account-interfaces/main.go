package main

import (
	"fmt"

	"github.com/otaxhu/go-methods-interfaces/bank-account-interfaces/interfaces"
	"github.com/otaxhu/go-methods-interfaces/bank-account-interfaces/models"
	"github.com/otaxhu/go-methods-interfaces/bank-account-interfaces/router"
)

func main() {
	router.InitializeRouter()
}

func printBalance(account interfaces.IBankAccounts) {
	fmt.Printf("Usted tiene %d en su cuenta de %s\n", account.GetBalance(), account.GetCurrency())
}

func printUserWithBalances(user *models.User) {
	for _, account := range user.GetBankAccounts() {
		fmt.Printf("este usuario %s tiene %d en su cuenta de %s\n", user.GetName(), account.GetBalance(), account.GetCurrency())
	}
}
