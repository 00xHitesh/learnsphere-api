package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/hitnarkhede/learnsphere-api/internal/config"
	"github.com/hitnarkhede/learnsphere-api/internal/middleware"
	"github.com/hitnarkhede/learnsphere-api/pkg/response"
)

type FeatureHandler struct {
	svc featureService
}

func NewFeatureHandler(svc featureService) *FeatureHandler {
	return &FeatureHandler{svc: svc}
}

func (h *FeatureHandler) RegisterRoutes(v1 *echo.Group, cfg *config.Config) {
	g := v1.Group("/features")
	g.GET("", h.List)
	g.PATCH("/:name/toggle", h.Toggle, middleware.JWT(cfg), middleware.RequireRole("admin"))
}

// List godoc
//
//	@Summary		List feature flags
//	@Description	Returns all feature flags and their enabled state
//	@Tags			Features
//	@Produce		json
//	@Success		200	{object}	response.Response{data=[]service.FeatureFlagResponse}
//	@Router			/features [get]
func (h *FeatureHandler) List(c echo.Context) error {
	flags, err := h.svc.List(c.Request().Context())
	if err != nil {
		return response.InternalError(c, "")
	}
	return response.OK(c, flags)
}

// Toggle godoc
//
//	@Summary		Toggle feature flag (admin)
//	@Description	Flip the enabled state of a named feature flag
//	@Tags			Features
//	@Produce		json
//	@Security		BearerAuth
//	@Param			name	path		string	true	"Feature flag name"
//	@Success		200		{object}	response.Response{data=service.FeatureFlagResponse}
//	@Failure		401		{object}	response.ErrorResponse
//	@Failure		403		{object}	response.ErrorResponse
//	@Router			/features/{name}/toggle [patch]
func (h *FeatureHandler) Toggle(c echo.Context) error {
	flags, err := h.svc.Toggle(c.Request().Context(), c.Param("name"))
	if err != nil {
		return response.InternalError(c, "")
	}
	return response.OK(c, flags)
}
