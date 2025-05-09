package onepass

import "errors"

// Common errors
var (
	ErrMissingToken     = errors.New("1password: token is required")
	ErrMissingServerURL = errors.New("1password: server URL is required")
	ErrPasswordNotFound = errors.New("1password: password not found in item")
	ErrFieldNotFound    = errors.New("1password: field not found in item")
	ErrEnvVarNotSet     = errors.New("environment variable not set")
)
