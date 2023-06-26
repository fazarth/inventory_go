package helpers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(ctx echo.Context, code int, data interface{}) error {
	response := Response{
		Status: "success",
		Code:   code,
		Data:   data,
	}
	return ctx.JSON(code, response)
}

func ErrorResponse(ctx echo.Context, code int, message string) error {
	response := Response{
		Status:  "error",
		Code:    code,
		Message: message,
	}
	return ctx.JSON(code, response)
}
