package validator

import (
	"errors"
)

var (
	MinErr  = errors.New("less than min")
	MaxErr  = errors.New("greater than max")
	ZeroErr = errors.New("zero value")
)
