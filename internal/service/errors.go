package service

import "errors"

var (
	ErrUserExists         = errors.New("user already exists")
	ErrNotFound           = errors.New("not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
)
