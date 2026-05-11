package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/hitnarkhede/learnsphere-api/internal/config"
	"github.com/hitnarkhede/learnsphere-api/internal/middleware"
	"github.com/hitnarkhede/learnsphere-api/pkg/response"
)

type WishlistHandler struct {
	svc wishlistService
}

func NewWishlistHandler(svc wishlistService) *WishlistHandler {
	return &WishlistHandler{svc: svc}
}

func (h *WishlistHandler) RegisterRoutes(v1 *echo.Group, cfg *config.Config) {
	g := v1.Group("/wishlist", middleware.JWT(cfg))
	g.GET("", h.ListByUser)
	g.POST("/toggle", h.Toggle)
}

// Request types

type WishlistToggleRequest struct {
	CourseID string `json:"courseId" example:"550e8400-e29b-41d4-a716-446655440000"`
}

// ListByUser godoc
//
//	@Summary		List wishlist
//	@Description	Returns all wishlist items for the authenticated user
//	@Tags			Wishlist
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{object}	response.Response{data=[]service.WishlistItemResponse}
//	@Failure		401	{object}	response.ErrorResponse
//	@Router			/wishlist [get]
func (h *WishlistHandler) ListByUser(c echo.Context) error {
	userID := c.Get("user_id").(string)
	items, err := h.svc.ListByUser(c.Request().Context(), userID)
	if err != nil {
		return response.InternalError(c, "")
	}
	return response.OK(c, items)
}

// Toggle godoc
//
//	@Summary		Toggle wishlist item
//	@Description	Add a course to the wishlist if not present, remove it if already present
//	@Tags			Wishlist
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			body	body		WishlistToggleRequest	true	"Course to toggle"
//	@Success		200		{object}	response.Response{data=object{wishlisted=bool}}
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		401		{object}	response.ErrorResponse
//	@Router			/wishlist/toggle [post]
func (h *WishlistHandler) Toggle(c echo.Context) error {
	var body struct {
		CourseID string `json:"courseId" validate:"required"`
	}
	if err := bindAndValidate(c, &body); err != nil {
		return err
	}
	userID := c.Get("user_id").(string)
	added, err := h.svc.Toggle(c.Request().Context(), userID, body.CourseID)
	if err != nil {
		return response.InternalError(c, "")
	}
	return response.OK(c, map[string]bool{"wishlisted": added})
}
