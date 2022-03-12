package util

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

//validate a given date which has validation rule(s) in go-playground validator schema link: https://github.com/go-playground/validator
func Validate(data interface{}) (string, error) {
	// declare validator
	var validate = validator.New()
	// declare string builder : error message
	var sb strings.Builder
	err := validate.Struct(data)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		//building error message
		for _, err := range err.(validator.ValidationErrors) {
			sb.WriteString(err.Field() + ":" + err.Tag() + ",")
			// other options for error message are:
			// fields.PushBack(fielderror)
			// fmt.Println(fielderror)
			// fmt.Println(err.Field())
			// fmt.Println(err.StructNamespace())
			// fmt.Println(err.StructField())
			// fmt.Println(err.Tag())
			// fmt.Println(err.ActualTag())
			// fmt.Println(err.Kind())
			// fmt.Println(err.Type())
			// fmt.Println(err.Value())
			// fmt.Println(err.Param())

		}

		//invalid data: from here you can create your own error messages in whatever language you wish
		return sb.String(), errors.New("validation failed")
	}
	//valid data
	return "", nil
}
