package main

import (
	"fmt"

	"github.com/otaxhu/go-methods-interfaces/bank-account-interfaces/interfaces"
	"github.com/otaxhu/go-methods-interfaces/bank-account-interfaces/models"
)

func main() {
	//router.InitializeRouter()
	cuentaBolivares := models.NewBolivaresAccount()
	cuentaDolares := models.NewUSDaccount()
	printBalance(cuentaDolares)
	printBalance(cuentaBolivares)
}

func printBalance(account interfaces.IBankAccounts) {
	fmt.Printf("Usted tiene %d en su cuenta de %s\n", account.GetBalance(), account.GetCurrency())
}
