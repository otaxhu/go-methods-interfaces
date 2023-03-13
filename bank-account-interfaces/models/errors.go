package models

import "errors"

var (
	errInsufficientFunds       = errors.New("fondos insuficientes")
	errInvalidAmount           = errors.New("monto invalido")
	errCurrencyAccountNotFound = errors.New("no se encontró una cuenta con la moneda especificada")
)
