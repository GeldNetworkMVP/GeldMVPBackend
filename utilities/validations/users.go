package validations

import (
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
	"github.com/go-playground/validator/v10"
)

func ValidateUsers(e model.AppUser) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}
