package validation

import (
	"errors"

	"github.com/YamaguchiKoki/react-go-todo/adapter/validator"
)

var (
	errInvalidInstance = errors.New("Invalid validator instance")
)

const (
	InstanceGoPlayground int = iota
)

func NewValidatorFactory(instance int) (validator.Validator, error) {
	switch instance {
	case InstanceGoPlayground:
		return NewGoPlayground()
	default:
		return nil, errInvalidInstance
	}
}