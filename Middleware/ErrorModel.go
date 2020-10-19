package Middleware

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type ErrorModel struct{
	StatusCode int `json:"Status code"`
	ErrorCode int `json:"Error code"`
	Message string `json:"Message"`
}

func (e *ErrorModel) Error() string {
	return strconv.Itoa(e.StatusCode)
}


func (e *ErrorModel) ResponseBody() ([]byte, error) {
	body, err := json.Marshal(e)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing response body: %v", err)
	}
	return body, nil
}

// ResponseHeaders returns http status code and headers.
func (e *ErrorModel) ResponseHeaders() (int, map[string]string) {
	return e.StatusCode, map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}
}

func newErrorModel(statusCode int, errorCode int, message string) error {
	return &ErrorModel{
		StatusCode:     statusCode,
		ErrorCode:       errorCode,
		Message:      message,
	}
}
