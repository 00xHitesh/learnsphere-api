package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/hitnarkhede/learnsphere-api/internal/config"
	"github.com/hitnarkhede/learnsphere-api/internal/middleware"
	"github.com/hitnarkhede/learnsphere-api/internal/service"
	"github.com/hitnarkhede/learnsphere-api/pkg/response"
)

type CourseHandler struct {
	svc courseService
}

func NewCourseHandler(svc courseService) *CourseHandler {
	return &CourseHandler{svc: svc}
}

func (h *CourseHandler) RegisterRoutes(v1 *echo.Group, cfg *config.Config) {
	g := v1.Group("/courses")
	g.GET("", h.List)
	g.GET("/:slug", h.GetBySlug)

	admin := g.Group("", middleware.JWT(cfg), middleware.RequireRole("admin"))
	admin.GET("/all", h.ListAll)
	admin.PUT("", h.Upsert)
	admin.DELETE("/:id", h.Delete)
	admin.PATCH("/:id/status", h.SetStatus)
}

// Request types

type UpsertCourseRequest struct {
	ID            string   `json:"id" example:""`
	Title         string   `json:"title" example:"Foundations of TypeScript"`
	Slug          string   `json:"slug" example:"foundations-of-typescript"`
	Description   string   `json:"description" example:"Move from JavaScript to confident TypeScript."`
	Summary       string   `json:"summary" example:"Master TypeScript fundamentals."`
	Category      string   `json:"category" example:"Programming"`
	Level         string   `json:"level" example:"Beginner"`
	ListPrice     float64  `json:"listPrice" example:"49"`
	SalePrice     *float64 `json:"salePrice" example:"29"`
	Instructor    string   `json:"instructor" example:"Kenji Romero"`
	InstructorBio string   `json:"instructorBio" example:"Senior engineer."`
	Cover         string   `json:"cover" example:"TS"`
	Status        string   `json:"status" example:"published"`
	Tags          []string `json:"tags" example:"typescript,programming"`
}

type SetCourseStatusRequest struct {
	Status string `json:"status" example:"published"`
}

// List godoc
//
//	@Summary		List published courses
//	@Description	Returns all courses with status=published, each with tags
//	@Tags			Courses
//	@Produce		json
//	@Success		200	{object}	response.Response{data=[]service.CourseResponse}
//	@Router			/courses [get]
func (h *CourseHandler) List(c echo.Context) error {
	courses, err := h.svc.ListPublished(c.Request().Context())
	if err != nil {
		return response.InternalError(c, "")
	}
	return response.OK(c, courses)
}

// ListAll godoc
//
//	@Summary		List all courses (admin)
//	@Description	Returns every course regardless of status
//	@Tags			Courses
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{object}	response.Response{data=[]service.CourseResponse}
//	@Failure		401	{object}	response.ErrorResponse
//	@Failure		403	{object}	response.ErrorResponse
//	@Router			/courses/all [get]
func (h *CourseHandler) ListAll(c echo.Context) error {
	courses, err := h.svc.ListAll(c.Request().Context())
	if err != nil {
		return response.InternalError(c, "")
	}
	return response.OK(c, courses)
}

// GetBySlug godoc
//
//	@Summary		Get course by slug
//	@Description	Returns a published course with full detail: lessons, reviews, quiz
//	@Tags			Courses
//	@Produce		json
//	@Param			slug	path		string	true	"Course slug"	example(foundations-of-typescript)
//	@Success		200		{object}	response.Response{data=service.CourseResponse}
//	@Failure		404		{object}	response.ErrorResponse
//	@Router			/courses/{slug} [get]
func (h *CourseHandler) GetBySlug(c echo.Context) error {
	course, err := h.svc.GetBySlug(c.Request().Context(), c.Param("slug"))
	if err != nil {
		return response.InternalError(c, "")
	}
	if course == nil {
		return response.NotFound(c, "course not found")
	}
	return response.OK(c, course)
}

// Upsert godoc
//
//	@Summary		Create or update a course (admin)
//	@Description	If id is empty a new course is created; otherwise the existing course is updated
//	@Tags			Courses
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			body	body		UpsertCourseRequest	true	"Course payload"
//	@Success		200		{object}	response.Response{data=service.CourseResponse}
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		401		{object}	response.ErrorResponse
//	@Failure		403		{object}	response.ErrorResponse
//	@Router			/courses [put]
func (h *CourseHandler) Upsert(c echo.Context) error {
	var body struct {
		ID            string   `json:"id"`
		Title         string   `json:"title" validate:"required"`
		Slug          string   `json:"slug" validate:"required"`
		Description   string   `json:"description" validate:"required"`
		Summary       string   `json:"summary" validate:"required"`
		Category      string   `json:"category" validate:"required"`
		Level         string   `json:"level" validate:"required"`
		ListPrice     float64  `json:"listPrice"`
		SalePrice     *float64 `json:"salePrice"`
		Instructor    string   `json:"instructor" validate:"required"`
		InstructorBio string   `json:"instructorBio" validate:"required"`
		Cover         string   `json:"cover" validate:"required"`
		Status        string   `json:"status" validate:"required"`
		Tags          []string `json:"tags"`
	}
	if err := bindAndValidate(c, &body); err != nil {
		return err
	}
	course, err := h.svc.Upsert(c.Request().Context(), service.UpsertCourseParams{
		ID:            body.ID,
		Title:         body.Title,
		Slug:          body.Slug,
		Description:   body.Description,
		Summary:       body.Summary,
		Category:      body.Category,
		Level:         body.Level,
		ListPrice:     body.ListPrice,
		SalePrice:     body.SalePrice,
		Instructor:    body.Instructor,
		InstructorBio: body.InstructorBio,
		Cover:         body.Cover,
		Status:        body.Status,
		Tags:          body.Tags,
	})
	if err != nil {
		return response.InternalError(c, err.Error())
	}
	return response.OK(c, course)
}

// Delete godoc
//
//	@Summary		Delete a course (admin)
//	@Description	Permanently deletes a course and its associated enrollments
//	@Tags			Courses
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path	string	true	"Course UUID"
//	@Success		204	"No content"
//	@Failure		401	{object}	response.ErrorResponse
//	@Failure		403	{object}	response.ErrorResponse
//	@Router			/courses/{id} [delete]
func (h *CourseHandler) Delete(c echo.Context) error {
	if err := h.svc.Delete(c.Request().Context(), c.Param("id")); err != nil {
		return response.InternalError(c, "")
	}
	return noContent(c)
}

// SetStatus godoc
//
//	@Summary		Set course status (admin)
//	@Description	Update the publication status of a course
//	@Tags			Courses
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id		path		string				true	"Course UUID"
//	@Param			body	body		SetCourseStatusRequest	true	"New status"
//	@Success		200		{object}	response.Response{data=service.CourseResponse}
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		401		{object}	response.ErrorResponse
//	@Failure		403		{object}	response.ErrorResponse
//	@Failure		404		{object}	response.ErrorResponse
//	@Router			/courses/{id}/status [patch]
func (h *CourseHandler) SetStatus(c echo.Context) error {
	var body struct {
		Status string `json:"status" validate:"required"`
	}
	if err := bindAndValidate(c, &body); err != nil {
		return err
	}
	course, err := h.svc.SetStatus(c.Request().Context(), c.Param("id"), body.Status)
	if err != nil {
		return response.InternalError(c, "")
	}
	if course == nil {
		return response.NotFound(c, "course not found")
	}
	return response.OK(c, course)
}
