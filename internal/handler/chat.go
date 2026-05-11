package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/hitnarkhede/learnsphere-api/internal/config"
	"github.com/hitnarkhede/learnsphere-api/internal/middleware"
	"github.com/hitnarkhede/learnsphere-api/pkg/response"
)

type ChatHandler struct {
	svc chatService
}

func NewChatHandler(svc chatService) *ChatHandler {
	return &ChatHandler{svc: svc}
}

func (h *ChatHandler) RegisterRoutes(v1 *echo.Group, cfg *config.Config) {
	g := v1.Group("/chat", middleware.JWT(cfg))
	g.GET("", h.List)
	g.POST("/messages", h.Append)
	g.DELETE("", h.Reset)
}

// Request types

type AppendMessageRequest struct {
	Role     string  `json:"role" example:"user"`
	Content  string  `json:"content" example:"What is TypeScript?"`
	CourseID *string `json:"courseId" example:"550e8400-e29b-41d4-a716-446655440000"`
}

// List godoc
//
//	@Summary		List chat messages
//	@Description	Returns chat history for the authenticated user, optionally filtered by course
//	@Tags			Chat
//	@Produce		json
//	@Security		BearerAuth
//	@Param			courseId	query		string	false	"Filter by course UUID"
//	@Success		200			{object}	response.Response{data=[]service.ChatMessageResponse}
//	@Failure		401			{object}	response.ErrorResponse
//	@Router			/chat [get]
func (h *ChatHandler) List(c echo.Context) error {
	userID := c.Get("user_id").(string)
	courseID := queryCourseID(c)
	msgs, err := h.svc.List(c.Request().Context(), userID, courseID)
	if err != nil {
		return response.InternalError(c, "")
	}
	return response.OK(c, msgs)
}

// Append godoc
//
//	@Summary		Append chat message
//	@Description	Add a new message (user or assistant) to the chat history
//	@Tags			Chat
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			body	body		AppendMessageRequest	true	"Message to append"
//	@Success		200		{object}	response.Response{data=service.ChatMessageResponse}
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		401		{object}	response.ErrorResponse
//	@Router			/chat/messages [post]
func (h *ChatHandler) Append(c echo.Context) error {
	var body struct {
		Role     string  `json:"role" validate:"required,oneof=user assistant"`
		Content  string  `json:"content" validate:"required"`
		CourseID *string `json:"courseId"`
	}
	if err := bindAndValidate(c, &body); err != nil {
		return err
	}
	userID := c.Get("user_id").(string)
	msg, err := h.svc.Append(c.Request().Context(), userID, body.CourseID, body.Role, body.Content)
	if err != nil {
		return response.InternalError(c, "")
	}
	return response.OK(c, msg)
}

// Reset godoc
//
//	@Summary		Reset chat history
//	@Description	Delete all chat messages for the authenticated user, optionally scoped to a course
//	@Tags			Chat
//	@Produce		json
//	@Security		BearerAuth
//	@Param			courseId	query	string	false	"Scope reset to a course UUID"
//	@Success		204			"No content"
//	@Failure		401			{object}	response.ErrorResponse
//	@Router			/chat [delete]
func (h *ChatHandler) Reset(c echo.Context) error {
	userID := c.Get("user_id").(string)
	courseID := queryCourseID(c)
	if err := h.svc.Reset(c.Request().Context(), userID, courseID); err != nil {
		return response.InternalError(c, "")
	}
	return noContent(c)
}

func queryCourseID(c echo.Context) *string {
	val := c.QueryParam("courseId")
	if val == "" {
		return nil
	}
	return &val
}
