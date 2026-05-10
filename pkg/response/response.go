package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Errors  any    `json:"errors,omitempty"`
}

func OK(c echo.Context, data any) error {
	return c.JSON(http.StatusOK, Response{Success: true, Data: data})
}

func Created(c echo.Context, data any) error {
	return c.JSON(http.StatusCreated, Response{Success: true, Data: data})
}

func NoContent(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

func BadRequest(c echo.Context, message string, errs ...any) error {
	r := ErrorResponse{Success: false, Message: message}
	if len(errs) > 0 {
		r.Errors = errs[0]
	}
	return c.JSON(http.StatusBadRequest, r)
}

func Unauthorized(c echo.Context, message string) error {
	return c.JSON(http.StatusUnauthorized, ErrorResponse{Success: false, Message: message})
}

func Forbidden(c echo.Context, message string) error {
	return c.JSON(http.StatusForbidden, ErrorResponse{Success: false, Message: message})
}

func NotFound(c echo.Context, message string) error {
	return c.JSON(http.StatusNotFound, ErrorResponse{Success: false, Message: message})
}

func Conflict(c echo.Context, message string) error {
	return c.JSON(http.StatusConflict, ErrorResponse{Success: false, Message: message})
}

func InternalError(c echo.Context, message string) error {
	return c.JSON(http.StatusInternalServerError, ErrorResponse{Success: false, Message: message})
}
