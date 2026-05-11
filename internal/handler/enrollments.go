package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/hitnarkhede/learnsphere-api/internal/config"
	"github.com/hitnarkhede/learnsphere-api/internal/middleware"
	"github.com/hitnarkhede/learnsphere-api/pkg/response"
)

type EnrollmentHandler struct {
	svc enrollmentService
}

func NewEnrollmentHandler(svc enrollmentService) *EnrollmentHandler {
	return &EnrollmentHandler{svc: svc}
}

func (h *EnrollmentHandler) RegisterRoutes(v1 *echo.Group, cfg *config.Config) {
	g := v1.Group("/enrollments", middleware.JWT(cfg))
	g.GET("", h.ListByUser)
	g.GET("/:courseId", h.Get)
	g.POST("", h.Enroll)
	g.PATCH("/:courseId/lessons/:lessonId/toggle", h.ToggleLessonComplete)
	g.PATCH("/:courseId/time", h.RecordTimeSpent)
	g.PATCH("/:courseId/quiz/:quizId/pass", h.MarkQuizPassed)
}

// Request types

type EnrollRequest struct {
	CourseID string `json:"courseId" example:"550e8400-e29b-41d4-a716-446655440000"`
}

type RecordTimeRequest struct {
	Seconds int32 `json:"seconds" example:"120"`
}

// ListByUser godoc
//
//	@Summary		List my enrollments
//	@Description	Returns all course enrollments for the authenticated user
//	@Tags			Enrollments
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{object}	response.Response{data=[]service.EnrollmentResponse}
//	@Failure		401	{object}	response.ErrorResponse
//	@Router			/enrollments [get]
func (h *EnrollmentHandler) ListByUser(c echo.Context) error {
	userID := c.Get("user_id").(string)
	enrollments, err := h.svc.ListByUser(c.Request().Context(), userID)
	if err != nil {
		return response.InternalError(c, "")
	}
	return response.OK(c, enrollments)
}

// Get godoc
//
//	@Summary		Get enrollment
//	@Description	Returns the authenticated user's enrollment for a specific course
//	@Tags			Enrollments
//	@Produce		json
//	@Security		BearerAuth
//	@Param			courseId	path		string	true	"Course UUID"
//	@Success		200			{object}	response.Response{data=service.EnrollmentResponse}
//	@Failure		401			{object}	response.ErrorResponse
//	@Failure		404			{object}	response.ErrorResponse
//	@Router			/enrollments/{courseId} [get]
func (h *EnrollmentHandler) Get(c echo.Context) error {
	userID := c.Get("user_id").(string)
	e, err := h.svc.Get(c.Request().Context(), userID, c.Param("courseId"))
	if err != nil {
		return response.InternalError(c, "")
	}
	if e == nil {
		return response.NotFound(c, "enrollment not found")
	}
	return response.OK(c, e)
}

// Enroll godoc
//
//	@Summary		Enroll in a course
//	@Description	Enroll the authenticated user in a course (free enroll; use /transactions for paid)
//	@Tags			Enrollments
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			body	body		EnrollRequest	true	"Course to enroll in"
//	@Success		200		{object}	response.Response{data=service.EnrollmentResponse}
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		401		{object}	response.ErrorResponse
//	@Router			/enrollments [post]
func (h *EnrollmentHandler) Enroll(c echo.Context) error {
	var body struct {
		CourseID string `json:"courseId" validate:"required"`
	}
	if err := bindAndValidate(c, &body); err != nil {
		return err
	}
	userID := c.Get("user_id").(string)
	e, err := h.svc.Enroll(c.Request().Context(), userID, body.CourseID)
	if err != nil {
		return response.InternalError(c, "")
	}
	return response.OK(c, e)
}

// ToggleLessonComplete godoc
//
//	@Summary		Toggle lesson completion
//	@Description	Mark a lesson as complete or incomplete and return the updated enrollment
//	@Tags			Enrollments
//	@Produce		json
//	@Security		BearerAuth
//	@Param			courseId	path		string	true	"Course UUID"
//	@Param			lessonId	path		string	true	"Lesson UUID"
//	@Success		200			{object}	response.Response{data=service.EnrollmentResponse}
//	@Failure		400			{object}	response.ErrorResponse
//	@Failure		401			{object}	response.ErrorResponse
//	@Failure		404			{object}	response.ErrorResponse
//	@Router			/enrollments/{courseId}/lessons/{lessonId}/toggle [patch]
func (h *EnrollmentHandler) ToggleLessonComplete(c echo.Context) error {
	userID := c.Get("user_id").(string)
	e, err := h.svc.ToggleLessonComplete(c.Request().Context(), userID, c.Param("courseId"), c.Param("lessonId"))
	if err != nil {
		return response.BadRequest(c, err.Error())
	}
	if e == nil {
		return response.NotFound(c, "enrollment not found")
	}
	return response.OK(c, e)
}

// RecordTimeSpent godoc
//
//	@Summary		Record time spent
//	@Description	Add seconds to the total time spent on a course enrollment
//	@Tags			Enrollments
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			courseId	path		string				true	"Course UUID"
//	@Param			body		body		RecordTimeRequest	true	"Seconds to add"
//	@Success		200			{object}	response.Response{data=service.EnrollmentResponse}
//	@Failure		400			{object}	response.ErrorResponse
//	@Failure		401			{object}	response.ErrorResponse
//	@Failure		404			{object}	response.ErrorResponse
//	@Router			/enrollments/{courseId}/time [patch]
func (h *EnrollmentHandler) RecordTimeSpent(c echo.Context) error {
	var body struct {
		Seconds int32 `json:"seconds" validate:"required,min=1"`
	}
	if err := bindAndValidate(c, &body); err != nil {
		return err
	}
	userID := c.Get("user_id").(string)
	e, err := h.svc.RecordTimeSpent(c.Request().Context(), userID, c.Param("courseId"), body.Seconds)
	if err != nil {
		return response.InternalError(c, "")
	}
	if e == nil {
		return response.NotFound(c, "enrollment not found")
	}
	return response.OK(c, e)
}

// MarkQuizPassed godoc
//
//	@Summary		Mark quiz passed
//	@Description	Record that the user has passed a quiz for a course enrollment
//	@Tags			Enrollments
//	@Produce		json
//	@Security		BearerAuth
//	@Param			courseId	path		string	true	"Course UUID"
//	@Param			quizId		path		string	true	"Quiz UUID"
//	@Success		200			{object}	response.Response{data=service.EnrollmentResponse}
//	@Failure		401			{object}	response.ErrorResponse
//	@Failure		404			{object}	response.ErrorResponse
//	@Router			/enrollments/{courseId}/quiz/{quizId}/pass [patch]
func (h *EnrollmentHandler) MarkQuizPassed(c echo.Context) error {
	userID := c.Get("user_id").(string)
	e, err := h.svc.MarkQuizPassed(c.Request().Context(), userID, c.Param("courseId"), c.Param("quizId"))
	if err != nil {
		return response.InternalError(c, "")
	}
	if e == nil {
		return response.NotFound(c, "enrollment not found")
	}
	return response.OK(c, e)
}
