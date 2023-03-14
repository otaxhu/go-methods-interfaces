package models

import "github.com/golang-jwt/jwt"

type appClaims struct {
	userId string
	jwt.StandardClaims
}

func NewAppClaims() *appClaims {
	return &appClaims{}
}

func (a *appClaims) SetUserId(id string) *appClaims {
	a.userId = id
	return a
}
