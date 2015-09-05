package validator

import (
	"fmt"
)

var DefaultValidator = NewValidator()

// NewValidator return the default validator
func NewValidator() *Validator {
	return &Validator{
		ErrMsgLoaction: "",
		ValidFunc:      make(map[string]validFunc),
	}
}

type validFunc func() error

type Validator struct {
	ErrMsgLoaction string
	ValidFunc      map[string]validFunc
}

// Validate the way to validate struct
func (v *Validator) Validate() error {

	return nil
}

// Validate use the default Validator to validate
func Validate() error {
	return DefaultValidator.Validate()
}

// load if the errMsgLoaction is given,the error message will be coverd.
func (v *Validator) load() error {
	if v.ErrMsgLoaction != "" {
		// TODO:加载错误消息，格式待定
	}

	return nil
}

// SetErrMsgLoaction
func (v *Validator) SetErrMsgLoaction(path string) {
	v.ErrMsgLoaction = path
}

// SetValidFunc register the function you want to valid
func (v *Validator) SetValidFunc(name string, f func() error) error {
	if name == "" {
		return fmt.Errorf("%s", "the function name should not be empty")
	}

	// delete the function
	if f == nil {
		delete(name, v.ValidFunc)
	}

	v.ValidFunc[name] = f

	return nil
}
