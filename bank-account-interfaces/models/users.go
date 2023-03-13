package models

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type User struct {
	Id       string        `json:"id" gorm:"primaryKey;not null;type:varchar(255)"`
	Name     string        `json:"name" gorm:"unique;not null;type:varchar(255)"`
	Accounts []BankAccount `json:"accounts" gorm:"foreignKey:user_id"`
}

func NewUser() *User {
	return &User{}
}

func (u *User) GetId() string {
	return u.Id
}

// This method already has an uuid.NewString()
func (u *User) NewUUID() *User {
	u.Id = uuid.NewString()
	return u
}

func (u *User) SetName(name string) *User {
	u.Name = name
	return u
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) NewBankAccount(currency string) *User {
	for _, crcy := range currencies {
		if crcy == currency {
			u.Accounts = append(u.Accounts, BankAccount{UserId: u.Id, Currency: crcy})
			return u
		}
	}
	//u.Error = errCurrencyAccountNotFound
	return u
}

func (u *User) GetBankAccounts() []BankAccount {
	return u.Accounts
}

func (u *User) Deposit(amount int, currency string) error {
	for _, account := range u.Accounts {
		if account.GetCurrency() == currency {
			if err := account.Deposit(amount); err != nil {
				return err
			}
			return nil
		}
	}
	return errCurrencyAccountNotFound
}

func (u *User) Withdraw(amount int, currency string) error {
	for _, account := range u.Accounts {
		if account.GetCurrency() == currency {
			if err := account.Withdraw(amount); err != nil {
				return err
			}
			return nil
		}
	}
	return errCurrencyAccountNotFound
}

func (u *User) ParseUserForm(r *http.Request) (*User, error) {
	if err := r.ParseMultipartForm(10 >> 20); err != nil {
		return nil, err
	}
	u.Name = strings.TrimSpace(r.FormValue("name"))
	if len(u.Name) < 8 {
		return nil, errInvalidUser
	}
	return u, nil
}

func (u *User) ParseClaims(tokenString string) *User {
	token, err := jwt.ParseWithClaims(tokenString, &appClaims{}, func(t *jwt.Token) (interface{}, error) { return os.Getenv("jwt_secret"), nil })
	if err != nil {
		//u.Error = err
		return u
	}
	claims, ok := token.Claims.(*appClaims)
	if !ok || !token.Valid {
		//u.Error = errInvalidToken
		return u
	}
	u.Id = claims.userId
	return u
}
