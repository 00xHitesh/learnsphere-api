package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/hitnarkhede/learnsphere-api/internal/config"
	"github.com/hitnarkhede/learnsphere-api/internal/middleware"
	"github.com/hitnarkhede/learnsphere-api/pkg/response"
)

type UserHandler struct {
	svc userService
}

func NewUserHandler(svc userService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) RegisterRoutes(v1 *echo.Group, cfg *config.Config) {
	g := v1.Group("/users", middleware.JWT(cfg), middleware.RequireRole("admin"))
	g.GET("", h.List)
	g.DELETE("/:id", h.Delete)
}

// List godoc
//
//	@Summary		List users (admin)
//	@Description	Returns all registered users
//	@Tags			Users
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{object}	response.Response{data=[]service.UserResponse}
//	@Failure		401	{object}	response.ErrorResponse
//	@Failure		403	{object}	response.ErrorResponse
//	@Router			/users [get]
func (h *UserHandler) List(c echo.Context) error {
	users, err := h.svc.List(c.Request().Context())
	if err != nil {
		return response.InternalError(c, "")
	}
	return response.OK(c, users)
}

// Delete godoc
//
//	@Summary		Delete user (admin)
//	@Description	Permanently delete a user account
//	@Tags			Users
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path	string	true	"User UUID"
//	@Success		204	"No content"
//	@Failure		401	{object}	response.ErrorResponse
//	@Failure		403	{object}	response.ErrorResponse
//	@Router			/users/{id} [delete]
func (h *UserHandler) Delete(c echo.Context) error {
	if err := h.svc.Delete(c.Request().Context(), c.Param("id")); err != nil {
		return response.InternalError(c, "")
	}
	return noContent(c)
}
