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
	code := http.StatusInternalServerError
	if msg != "" {
		resp["message"] = msg
	}

	switch true {
	case strings.Contains(msg, "server"):
		code = http.StatusInternalServerError
	case strings.Contains(msg, "format"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "kurang"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "role/status tidak valid"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "tidak ditemukan"):
		code = http.StatusNotFound
	case strings.Contains(msg, "not found"):
		code = http.StatusNotFound
	case strings.Contains(msg, "conflict"):
		code = http.StatusConflict
	case strings.Contains(msg, "duplicated"):
		code = http.StatusConflict
	case strings.Contains(msg, "syntax"):
		code = http.StatusNotFound
		resp["message"] = "not found"
	case strings.Contains(msg, "input invalid"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "input value"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "validation"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "unmarshal"):
		resp["message"] = "failed to unmarshal json"
		code = http.StatusBadRequest
	case strings.Contains(msg, "upload"):
		code = http.StatusInternalServerError
	case strings.Contains(msg, "denied"):
		code = http.StatusUnauthorized
	case strings.Contains(msg, "jwt"):
		msg = "access is denied due to invalid credential"
		code = http.StatusUnauthorized
	}

	return code, resp
}

func ErrorResponse(msg string) interface{} {
	resp := map[string]interface{}{}
	resp["message"] = msg

	return resp
}

type PaginationResponse struct {
	Page        int `json:"page"`
	Limit       int `json:"limit"`
	Offset      int `json:"offset"`
	TotalRecord int `json:"total_record"`
	TotalPage   int `json:"total_page"`
}

type WithPagination struct {
	Pagination PaginationResponse `json:"pagination"`
	Data       interface{}        `json:"data"`
	Message    string             `json:"message"`
}
