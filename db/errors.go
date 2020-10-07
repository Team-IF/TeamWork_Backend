package db

import "errors"

var (
	ErrAlreadyVerified  = errors.New("already verified")
	ErrExpired          = errors.New("expired")
	ErrPasswordNotMatch = errors.New("password not match")
)
