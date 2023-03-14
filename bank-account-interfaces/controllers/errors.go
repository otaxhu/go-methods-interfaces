package controllers

import "errors"

var (
	errUnauthorized = errors.New("unauthorized")
	errRegistered   = errors.New("already registered")
)
