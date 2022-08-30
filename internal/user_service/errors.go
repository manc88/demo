package userservice

import "errors"

var (
	ErrService = errors.New("service error")
	ErrStorage = errors.New("storage error")
)
