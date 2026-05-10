package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/hitnarkhede/learnsphere-api/pkg/response"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Check godoc
//
//	@Summary		Health check
//	@Description	Returns service health status
//	@Tags			Health
//	@Produce		json
//	@Success		200	{object}	response.Response{data=object{status=string}}
//	@Router			/health [get]
func (h *HealthHandler) Check(c echo.Context) error {
	return response.OK(c, map[string]string{"status": "ok"})
}
