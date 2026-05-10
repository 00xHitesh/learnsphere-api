package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/hitnarkhede/learnsphere-api/internal/config"
	"github.com/hitnarkhede/learnsphere-api/internal/middleware"
	"github.com/hitnarkhede/learnsphere-api/internal/service"
	"github.com/hitnarkhede/learnsphere-api/pkg/response"
)

type TransactionHandler struct {
	svc *service.TransactionService
}

func NewTransactionHandler(svc *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{svc: svc}
}

func (h *TransactionHandler) RegisterRoutes(v1 *echo.Group, cfg *config.Config) {
	g := v1.Group("/transactions", middleware.JWT(cfg))
	g.GET("", h.ListByUser)
	g.POST("", h.Purchase)

	g.GET("/all", wrapAdmin(cfg, h.ListAll))
}

func wrapAdmin(cfg *config.Config, fn echo.HandlerFunc) echo.HandlerFunc {
	return middleware.RequireRole("admin")(fn)
}

// Request types

type PurchaseRequest struct {
	CourseID string `json:"courseId" example:"550e8400-e29b-41d4-a716-446655440000"`
}

// ListByUser godoc
//
//	@Summary		List my transactions
//	@Description	Returns all purchase transactions for the authenticated user
//	@Tags			Transactions
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{object}	response.Response{data=[]service.TransactionResponse}
//	@Failure		401	{object}	response.ErrorResponse
//	@Router			/transactions [get]
func (h *TransactionHandler) ListByUser(c echo.Context) error {
	userID := c.Get("user_id").(string)
	txs, err := h.svc.ListByUser(c.Request().Context(), userID)
	if err != nil {
		return response.InternalError(c, "")
	}
	return response.OK(c, txs)
}

// ListAll godoc
//
//	@Summary		List all transactions (admin)
//	@Description	Returns all transactions across all users
//	@Tags			Transactions
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{object}	response.Response{data=[]service.TransactionResponse}
//	@Failure		401	{object}	response.ErrorResponse
//	@Failure		403	{object}	response.ErrorResponse
//	@Router			/transactions/all [get]
func (h *TransactionHandler) ListAll(c echo.Context) error {
	txs, err := h.svc.ListAll(c.Request().Context())
	if err != nil {
		return response.InternalError(c, "")
	}
	return response.OK(c, txs)
}

// Purchase godoc
//
//	@Summary		Purchase a course
//	@Description	Create a transaction and enrollment for the authenticated user
//	@Tags			Transactions
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			body	body		PurchaseRequest	true	"Course to purchase"
//	@Success		200		{object}	response.Response{data=object{transaction=service.TransactionResponse,enrollment=service.EnrollmentResponse}}
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		401		{object}	response.ErrorResponse
//	@Router			/transactions [post]
func (h *TransactionHandler) Purchase(c echo.Context) error {
	var body struct {
		CourseID string `json:"courseId" validate:"required"`
	}
	if err := bindAndValidate(c, &body); err != nil {
		return err
	}
	userID := c.Get("user_id").(string)
	tx, enrollment, err := h.svc.Purchase(c.Request().Context(), userID, body.CourseID)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}
	return response.OK(c, map[string]interface{}{"transaction": tx, "enrollment": enrollment})
}
