package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/hitnarkhede/learnsphere-api/pkg/response"
)

func bindAndValidate(c echo.Context, v interface{}) error {
	if err := c.Bind(v); err != nil {
		return response.BadRequest(c, "invalid request body")
	}
	if err := c.Validate(v); err != nil {
		return response.BadRequest(c, err.Error())
	}
	return nil
}

func noContent(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
