package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Standard response structure
func formatResponse(message string, data interface{}, status string, meta interface{}) map[string]interface{} {
	response := map[string]interface{}{
		"message": message,
		"data":    data,
		"status":  status, // Mengubah error menjadi status
	}

	if meta != nil {
		response["pagination"] = meta // Sesuai format contoh API
	}

	return response
}

// General Success Response
func ResponseFormatter(ctx echo.Context, code int, message string, body interface{}) error {
	response := formatResponse(message, body, "success", nil)
	return ctx.JSON(code, response)
}

// Success Response with Pagination Meta
func ResponseFormatterWithMeta(ctx echo.Context, code int, message string, body interface{}, meta interface{}) error {
	response := formatResponse(message, body, "success", meta)
	return ctx.JSON(code, response)
}

// General Error Response
func ResponseError(ctx echo.Context, code int, message string, err error) error {
	errorMessage := ""
	if err != nil {
		errorMessage = err.Error()
	}

	response := formatResponse(errorMessage, nil, "error", nil)
	//response["error_message"] = errorMessage // Tambahkan error_message

	return ctx.JSON(code, response)
}

// Validation Error Response
func ResponseErrValidation(ctx echo.Context, message string, errMap map[string]interface{}) error {
	response := formatResponse(message, nil, "error", nil)
	response["errors"] = errMap // Tambahkan daftar error

	return ctx.JSON(http.StatusBadRequest, response)
}

// Validation Error Response with Custom Code
func ResponseErrValidationWithCode(ctx echo.Context, message string, errMap map[string]interface{}, code int) error {
	response := formatResponse(message, nil, "error", nil)
	response["errors"] = errMap

	return ctx.JSON(code, response)
}
