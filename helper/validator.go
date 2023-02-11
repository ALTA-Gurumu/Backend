package helper

import (
	"errors"
	"fmt"
	"net/mail"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func ValidateRegisterRequest(nama string, email string, password string) error {
	if nama == "" {
		return errors.New("nama tidak boleh kosomg")
	}
	if email == "" {
		return errors.New("email tidak boleh kosomg")
	}
	if password == "" {
		return errors.New("password tidak boleh kosomg")
	}
	return nil
}

func IsGoodName(nama string) bool {
	return len(nama) >= 6
}

func IsStrongPassword(password string) bool {
	return len(password) >= 6
}

func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsStructEmpty(s interface{}) bool {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		if !reflect.DeepEqual(v.Field(i).Interface(), reflect.Zero(v.Field(i).Type()).Interface()) {
			return true
		}
	}

	return false
}

func ValidationErrorHandle(err error) string {
	var message string

	castedObject, ok := err.(validator.ValidationErrors)
	if ok {
		for _, v := range castedObject {
			switch v.Tag() {
			case "required":
				message = fmt.Sprintf("%s input value is required", v.Field())
			case "min":
				message = fmt.Sprintf("%s input value must be greater than %s character", v.Field(), v.Param())
			case "max":
				message = fmt.Sprintf("%s input value must be lower than %s character", v.Field(), v.Param())
			case "lte":
				message = fmt.Sprintf("%s input value must be below %s", v.Field(), v.Param())
			case "gte":
				message = fmt.Sprintf("%s input value must be above %s", v.Field(), v.Param())
			case "numeric":
				message = fmt.Sprintf("%s input value must be numeic", v.Field())
			case "url":
				message = fmt.Sprintf("%s input value must be am url", v.Field())
			case "email":
				message = fmt.Sprintf("%s input value must be an email", v.Field())
			case "password":
				message = fmt.Sprintf("%s input value must be filled", v.Field())
			}
		}
	}

	return message
}
