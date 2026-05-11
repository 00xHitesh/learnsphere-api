package handler

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
	"github.com/hitnarkhede/learnsphere-api/internal/config"
	"github.com/hitnarkhede/learnsphere-api/internal/middleware"
	"github.com/hitnarkhede/learnsphere-api/internal/service"
	"github.com/hitnarkhede/learnsphere-api/pkg/response"
)

type LiveSessionHandler struct {
	svc liveSessionService
}

func NewLiveSessionHandler(svc liveSessionService) *LiveSessionHandler {
	return &LiveSessionHandler{svc: svc}
}

func (h *LiveSessionHandler) RegisterRoutes(v1 *echo.Group, cfg *config.Config) {
	g := v1.Group("/live-sessions")
	g.GET("", h.List)

	admin := g.Group("", middleware.JWT(cfg), middleware.RequireRole("admin"))
	admin.GET("/all", h.ListAll)
	admin.PUT("", h.Upsert)
	admin.DELETE("/:id", h.Delete)
	admin.PATCH("/:id/status", h.SetStatus)
}

// Request types

type UpsertLiveSessionRequest struct {
	ID              string   `json:"id" example:""`
	Title           string   `json:"title" example:"TypeScript Deep Dive"`
	Instructor      string   `json:"instructor" example:"Kenji Romero"`
	Description     string   `json:"description" example:"Advanced TypeScript patterns."`
	ScheduledAt     string   `json:"scheduledAt" example:"2025-09-01T10:00:00Z"`
	DurationMinutes int32    `json:"durationMinutes" example:"90"`
	Topic           string   `json:"topic" example:"TypeScript"`
	JoinURL         string   `json:"joinUrl" example:"https://meet.example.com/abc123"`
	CourseID        *string  `json:"courseId" example:"550e8400-e29b-41d4-a716-446655440000"`
	RegisteredCount int32    `json:"registeredCount" example:"0"`
	MaxAttendees    int32    `json:"maxAttendees" example:"100"`
	ListPrice       *float64 `json:"listPrice" example:"49"`
	SalePrice       *float64 `json:"salePrice" example:"29"`
	Status          string   `json:"status" example:"published"`
}

type SetLiveSessionStatusRequest struct {
	Status string `json:"status" example:"published"`
}

// List godoc
//
//	@Summary		List published live sessions
//	@Description	Returns all live sessions with status=published
//	@Tags			Live Sessions
//	@Produce		json
//	@Success		200	{object}	response.Response{data=[]service.LiveSessionResponse}
//	@Router			/live-sessions [get]
func (h *LiveSessionHandler) List(c echo.Context) error {
	sessions, err := h.svc.ListPublished(c.Request().Context())
	if err != nil {
		return response.InternalError(c, "")
	}
	return response.OK(c, sessions)
}

// ListAll godoc
//
//	@Summary		List all live sessions (admin)
//	@Description	Returns every live session regardless of status
//	@Tags			Live Sessions
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{object}	response.Response{data=[]service.LiveSessionResponse}
//	@Failure		401	{object}	response.ErrorResponse
//	@Failure		403	{object}	response.ErrorResponse
//	@Router			/live-sessions/all [get]
func (h *LiveSessionHandler) ListAll(c echo.Context) error {
	sessions, err := h.svc.ListAll(c.Request().Context())
	if err != nil {
		return response.InternalError(c, "")
	}
	return response.OK(c, sessions)
}

// Upsert godoc
//
//	@Summary		Create or update a live session (admin)
//	@Description	If id is empty a new session is created; otherwise the existing session is updated
//	@Tags			Live Sessions
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			body	body		UpsertLiveSessionRequest	true	"Session payload"
//	@Success		200		{object}	response.Response{data=service.LiveSessionResponse}
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		401		{object}	response.ErrorResponse
//	@Failure		403		{object}	response.ErrorResponse
//	@Router			/live-sessions [put]
func (h *LiveSessionHandler) Upsert(c echo.Context) error {
	var body struct {
		ID              string   `json:"id"`
		Title           string   `json:"title" validate:"required"`
		Instructor      string   `json:"instructor" validate:"required"`
		Description     string   `json:"description"`
		ScheduledAt     string   `json:"scheduledAt" validate:"required"`
		DurationMinutes int32    `json:"durationMinutes" validate:"required,min=1"`
		Topic           string   `json:"topic"`
		JoinURL         string   `json:"joinUrl" validate:"required"`
		CourseID        *string  `json:"courseId"`
		RegisteredCount int32    `json:"registeredCount"`
		MaxAttendees    int32    `json:"maxAttendees" validate:"required,min=1"`
		ListPrice       *float64 `json:"listPrice"`
		SalePrice       *float64 `json:"salePrice"`
		Status          string   `json:"status" validate:"required"`
	}
	if err := bindAndValidate(c, &body); err != nil {
		return err
	}
	t, err := time.Parse(time.RFC3339, body.ScheduledAt)
	if err != nil {
		return response.BadRequest(c, "invalid scheduledAt format, use RFC3339")
	}
	scheduledAt := pgtype.Timestamptz{Time: t, Valid: true}

	session, err := h.svc.Upsert(c.Request().Context(), service.UpsertLiveSessionParams{
		ID:              body.ID,
		Title:           body.Title,
		Instructor:      body.Instructor,
		Description:     body.Description,
		ScheduledAt:     scheduledAt,
		DurationMinutes: body.DurationMinutes,
		Topic:           body.Topic,
		JoinURL:         body.JoinURL,
		CourseID:        body.CourseID,
		RegisteredCount: body.RegisteredCount,
		MaxAttendees:    body.MaxAttendees,
		ListPrice:       body.ListPrice,
		SalePrice:       body.SalePrice,
		Status:          body.Status,
	})
	if err != nil {
		return response.InternalError(c, err.Error())
	}
	return response.OK(c, session)
}

// Delete godoc
//
//	@Summary		Delete a live session (admin)
//	@Description	Permanently deletes a live session
//	@Tags			Live Sessions
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path	string	true	"Session UUID"
//	@Success		204	"No content"
//	@Failure		401	{object}	response.ErrorResponse
//	@Failure		403	{object}	response.ErrorResponse
//	@Router			/live-sessions/{id} [delete]
func (h *LiveSessionHandler) Delete(c echo.Context) error {
	if err := h.svc.Delete(c.Request().Context(), c.Param("id")); err != nil {
		return response.InternalError(c, "")
	}
	return noContent(c)
}

// SetStatus godoc
//
//	@Summary		Set live session status (admin)
//	@Description	Update the publication status of a live session
//	@Tags			Live Sessions
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id		path		string						true	"Session UUID"
//	@Param			body	body		SetLiveSessionStatusRequest	true	"New status"
//	@Success		200		{object}	response.Response{data=service.LiveSessionResponse}
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		401		{object}	response.ErrorResponse
//	@Failure		403		{object}	response.ErrorResponse
//	@Failure		404		{object}	response.ErrorResponse
//	@Router			/live-sessions/{id}/status [patch]
func (h *LiveSessionHandler) SetStatus(c echo.Context) error {
	var body struct {
		Status string `json:"status" validate:"required"`
	}
	if err := bindAndValidate(c, &body); err != nil {
		return err
	}
	session, err := h.svc.SetStatus(c.Request().Context(), c.Param("id"), body.Status)
	if err != nil {
		return response.InternalError(c, "")
	}
	if session == nil {
		return response.NotFound(c, "session not found")
	}
	return response.OK(c, session)
}
