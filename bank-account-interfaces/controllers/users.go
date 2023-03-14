package controllers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt"
	"github.com/otaxhu/go-methods-interfaces/bank-account-interfaces/database"
	"github.com/otaxhu/go-methods-interfaces/bank-account-interfaces/models"
	"gorm.io/gorm"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	user, err := models.NewUser().NewUUID().ParseUserForm(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tx, err := database.NewTransaction()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tx.Create(&user).Error; err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			http.Error(w, errRegistered.Error(), http.StatusConflict)
			tx.Rollback()
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		tx.Rollback()
		return
	}
	tx.Commit()
	w.WriteHeader(http.StatusCreated)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	user, err := models.NewUser().ParseUserForm(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	tx, err := database.NewTransaction()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tx.First(&user, "name = ?", user.GetName()).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, errUnauthorized.Error(), http.StatusUnauthorized)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	claims := models.NewAppClaims().SetUserId(user.GetId())
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(os.Getenv("jwt_secret"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tokenString)
	tx.Commit()
}
