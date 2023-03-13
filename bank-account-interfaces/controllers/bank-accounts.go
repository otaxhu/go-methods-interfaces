package controllers

import (
	"net/http"

	"github.com/otaxhu/go-methods-interfaces/bank-account-interfaces/models"
)

func UserDeposit(w http.ResponseWriter, r *http.Request) {
	user := models.NewUser().DecodeJson(r.Body)

}
