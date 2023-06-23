package config

import (
	"net/http"
)

// Response represents the response structure
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

// ErrorResponse represents the error response structure
type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

// SendSuccessResponse sends a success response
func SendSuccessResponse(c echo.Context, data interface{}, message string) error {
	response := Response{
		Success: true,
		Data:    data,
		Message: message,
	}
	return c.JSON(http.StatusOK, response)
}

// SendErrorResponse sends an error response
func SendErrorResponse(c echo.Context, err error, statusCode int) error {
	response := ErrorResponse{
		Success: false,
		Error:   err.Error(),
	}
	return c.JSON(statusCode, response)
}