package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/hitnarkhede/learnsphere-api/internal/config"
	"github.com/hitnarkhede/learnsphere-api/internal/middleware"
	"github.com/hitnarkhede/learnsphere-api/internal/service"
	"github.com/hitnarkhede/learnsphere-api/pkg/response"
)

type StatsHandler struct {
	svc *service.StatsService
}

func NewStatsHandler(svc *service.StatsService) *StatsHandler {
	return &StatsHandler{svc: svc}
}

func (h *StatsHandler) RegisterRoutes(v1 *echo.Group, cfg *config.Config) {
	g := v1.Group("/stats", middleware.JWT(cfg), middleware.RequireRole("admin"))
	g.GET("/overview", h.Overview)
	g.GET("/revenue", h.Revenue)
	g.GET("/active-users", h.ActiveUsers)
}

// Overview godoc
//
//	@Summary		Admin overview stats
//	@Description	Returns total users, courses, enrollments, and revenue
//	@Tags			Stats
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{object}	response.Response{data=service.OverviewResponse}
//	@Failure		401	{object}	response.ErrorResponse
//	@Failure		403	{object}	response.ErrorResponse
//	@Router			/stats/overview [get]
func (h *StatsHandler) Overview(c echo.Context) error {
	data, err := h.svc.Overview(c.Request().Context())
	if err != nil {
		return response.InternalError(c, "")
	}
	return response.OK(c, data)
}

// Revenue godoc
//
//	@Summary		Revenue time series
//	@Description	Returns revenue grouped by period (daily, weekly, monthly)
//	@Tags			Stats
//	@Produce		json
//	@Security		BearerAuth
//	@Param			period	query		string	false	"Grouping period: daily, weekly, monthly"	Enums(daily, weekly, monthly)	default(daily)
//	@Success		200		{object}	response.Response{data=[]service.PeriodPoint}
//	@Failure		401		{object}	response.ErrorResponse
//	@Failure		403		{object}	response.ErrorResponse
//	@Router			/stats/revenue [get]
func (h *StatsHandler) Revenue(c echo.Context) error {
	period := c.QueryParam("period")
	if period == "" {
		period = "daily"
	}
	data, err := h.svc.RevenueByPeriod(c.Request().Context(), period)
	if err != nil {
		return response.InternalError(c, "")
	}
	return response.OK(c, data)
}

// ActiveUsers godoc
//
//	@Summary		Active users time series
//	@Description	Returns distinct active user counts grouped by period (daily, weekly, monthly)
//	@Tags			Stats
//	@Produce		json
//	@Security		BearerAuth
//	@Param			period	query		string	false	"Grouping period: daily, weekly, monthly"	Enums(daily, weekly, monthly)	default(daily)
//	@Success		200		{object}	response.Response{data=[]service.PeriodPoint}
//	@Failure		401		{object}	response.ErrorResponse
//	@Failure		403		{object}	response.ErrorResponse
//	@Router			/stats/active-users [get]
func (h *StatsHandler) ActiveUsers(c echo.Context) error {
	period := c.QueryParam("period")
	if period == "" {
		period = "daily"
	}
	data, err := h.svc.ActiveUsersByPeriod(c.Request().Context(), period)
	if err != nil {
		return response.InternalError(c, "")
	}
	return response.OK(c, data)
}
