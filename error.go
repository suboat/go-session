package gosession

import "errors"

var (
	ErrNil     error = errors.New("gosession: DON`T supoort nil")
	ErrSession       = errors.New("gosession: No session")
	ErrKey           = errors.New("gosession: Key can't be empty")
)
