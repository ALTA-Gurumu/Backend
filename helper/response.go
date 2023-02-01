package helper

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator"
)

func PrintSuccessReponse(code int, message string, data ...interface{}) (int, interface{}) {
	resp := map[string]interface{}{}
	switch len(data) {
	case 1:
		resp["data"] = data[0]
	case 2:
		resp["token"] = data[1].(string)
		resp["data"] = data[0]
	}
	if message != "" {
		resp["message"] = message
	}

	return code, resp
}

func PrintErrorResponse(msg string) (int, interface{}) {
	resp := map[string]interface{}{}
	code := -1
	if msg != "" {
		resp["message"] = msg
	}

	if strings.Contains(msg, "server") {
		code = http.StatusInternalServerError
	} else if strings.Contains(msg, "format") {
		code = http.StatusBadRequest
	} else if strings.Contains(msg, "not found") {
		code = http.StatusNotFound
	}

	return code, resp
}

func ErrorResponse(msg string) interface{} {
	resp := map[string]interface{}{}
	resp["message"] = msg

	return resp
}

func ValidationErrorHandle(err error) string {
	messages := []string{}

	castedObject, ok := err.(validator.ValidationErrors)
	if ok {
		for _, v := range castedObject {
			switch v.Tag() {
			case "required":
				messages = append(messages, fmt.Sprintf("%s is required", v.Field()))
			case "min":
				messages = append(messages, fmt.Sprintf("%s value must be greater than %s character", v.Field(), v.Param()))
			case "max":
				messages = append(messages, fmt.Sprintf("%s value must be lower than %s character", v.Field(), v.Param()))
			case "lte":
				messages = append(messages, fmt.Sprintf("%s value must be below %s", v.Field(), v.Param()))
			case "gte":
				messages = append(messages, fmt.Sprintf("%s value must be above %s", v.Field(), v.Param()))
			case "numeric":
				messages = append(messages, fmt.Sprintf("%s value must be numeic", v.Field()))
			case "url":
				messages = append(messages, fmt.Sprintf("%s value must be am url", v.Field()))
			case "email":
				messages = append(messages, fmt.Sprintf("%s value must be an email", v.Field()))
			case "password":
				messages = append(messages, fmt.Sprintf("%s value must be filled", v.Field()))
			}
		}
	}

	msg := strings.Join(messages, ", ")

	return msg
}
