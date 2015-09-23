package msg

import (
	"errors"
)

// Errors
var (
	ErrUnknownAction = errors.New("Unknown Action")
)

const (
	ActionHandshake = iota
	ActionGet
	ActionSet
)

const (
	StatusSuccess = iota
	StatusClientError
	StatusServerError
	StatusKeyNotFound
)
