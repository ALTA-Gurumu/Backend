package helper

import (
	"net/http"
	"strings"
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
