package util

import "fmt"

//var validate = validator.New()
//
//func Validation(data interface{}) map[string]interface{} {
//
//	fmt.Println("hello")
//
//	en := en.New()
//	uni := ut.New(en, en)
//
//	trans, _ := uni.GetTranslator("en")
//	en_translations.RegisterDefaultTranslations(validate, trans)
//
//	var errorMap = map[string]interface{}{}
//	fmt.Println(errorMap)
//
//	err := validate.Struct(data)
//
//	if err != nil {
//
//		errs := err.(validator.ValidationErrors)
//
//		var errorList = make([]string,len(errs))
//
//		for _, e := range errs {
//			errorList = append(errorList, e.Translate(trans))
//		}
//
//		errorMap["error"] = errorList
//	}
//
//	if len(errorMap)>0 {
//		return errorMap
//	}
//
//	return nil
//}

func Test(data interface{}) {

	fmt.Println("sss", data)
}
