package handler

import (
	"github.com/Jeffail/gabs"
	"github.com/go-playground/validator"
)

func ValidationErrorInJson(err error) interface{} {
	validationErrors := err.(validator.ValidationErrors)
	validationErrFriendlyJson := gabs.New()
	for _, v := range validationErrors {
		inputVsExpected := []string{v.Value().(string), v.ActualTag()}
		validationErrFriendlyJson.SetP(inputVsExpected, v.Namespace())
	}

	return validationErrFriendlyJson.Data()
}
