package validations

import (
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
	"github.com/go-playground/validator/v10"
)

func ValidateMasterData(e model.MasterData) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}

func ValidateMasterDataCollection(e model.DataCollection) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}
