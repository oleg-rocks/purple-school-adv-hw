package req

import (
	"github.com/go-playground/validator/v10"
)

func IsValid[T any](payload T) error {
	validation := validator.New()
	err := validation.Struct(payload)
	if err != nil {
		return err
	}
	return nil
}