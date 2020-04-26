package util

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var validate = validator.New()

func Validation(data interface{}) map[string]interface{} {

	en := en.New()
	uni := ut.New(en, en)

	trans, _ := uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(validate, trans)

	var errorMap = make(map[string]interface{}, 1)

	err := validate.Struct(data)

	if err != nil {

		errs := err.(validator.ValidationErrors)

		var errorList = make([]string, 0)

		for _, e := range errs {
			fmt.Println(e.Translate(trans))
			errorList = append(errorList, e.Translate(trans))
		}

		errorMap["error"] = errorList
	}

	if len(errorMap) > 0 {
		return errorMap
	}

	return nil
}
