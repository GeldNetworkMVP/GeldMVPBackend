package validations

import (
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
	"github.com/go-playground/validator/v10"
)

func ValidateToken(e model.TokenPayload) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}
