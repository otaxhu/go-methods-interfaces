package models

import "errors"

var (
	errInsufficientFunds = errors.New("fondos insuficientes")
	errInvalidAmount     = errors.New("monto invalido")
)
