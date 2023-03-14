package models

import "errors"

var (
	errInsufficientFunds       = errors.New("fondos insuficientes")
	errInvalidAmount           = errors.New("monto invalido")
	errCurrencyAccountNotFound = errors.New("no se encontró una cuenta con la moneda especificada")
	errInvalidUser             = errors.New("usuario invalido")
	errInvalidToken            = errors.New("token invalido")
)

var currencies = []string{
	"Bitcoin", "Bolivares", "Dolares",
}
