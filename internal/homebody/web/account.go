package web

import (
	"errors"
)

var (
	ErrEmptyUser        = errors.New("empty user name")
	ErrEmptyAccountInfo = errors.New("empty account info")
	ErrEmptyDate        = errors.New("empty date")
)

var (
	setAccountPath    = "/account/set"
	getAccountPath    = "/account/get"
	updateAccountPath = "/account/update"
	deleteAccountPath = "/account/delete"
)
