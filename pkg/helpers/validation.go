package helpers

import (
	"github.com/asaskevich/govalidator"
	"github.com/fydhfzh/fp-4/pkg/errs"
)


func ValidateStruct(payload interface{}) errs.Errs {

	_, err := govalidator.ValidateStruct(payload)

	if err != nil {
		return errs.NewBadRequestError(err.Error())
	}

	return nil
}