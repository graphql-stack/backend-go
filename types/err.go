package types

import "errors"

var (
	ErrValidateError        = errors.New("INVALID_INPUT")
	ErrEmailOrPasswordError = errors.New("INVALID_USERNAME_OR_PASSWORD")
	ErrInternalError        = errors.New("INTERNAL_ERROR")
	ErrInvalidToken         = errors.New("INVALID_TOKEN")
)
