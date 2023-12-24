package handlers

import "github.com/labstack/echo/v4"

type commonResponse struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	ErrorData *ErrData    `json:"error,omitempty"`
}

type ErrData struct {
	ErrorMessage string `json:"error_message"`
}

func responseOK(c echo.Context, respCode int, message string, data interface{}) error {

	return c.JSON(respCode, commonResponse{
		Code:    respCode,
		Message: message,
		Data:    data,
	})
}

func responseError(c echo.Context, respCode int, message string, data interface{}, errMessage string) error {

	return c.JSON(respCode, commonResponse{
		Code:      respCode,
		Message:   message,
		Data:      data,
		ErrorData: &ErrData{ErrorMessage: errMessage},
	})
}
