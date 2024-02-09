package types

import (
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

func Required(fieldName string) *requiredRule {
	return &requiredRule{
		message: fmt.Sprintf("%s can not be empty", fieldName),
	}
}

type requiredRule struct {
	message string
}

// Validate checks if the given value is valid or not.
func (v *requiredRule) Validate(value interface{}) error {
	value, isNil := validation.Indirect(value)
	if isNil || validation.IsEmpty(value) {
		return errors.New(v.message)
	}
	return nil
}

// Error sets the error message for the rule.
func (v *requiredRule) Error(message string) *requiredRule {
	return &requiredRule{
		message: message,
	}
}
