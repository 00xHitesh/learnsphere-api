package handler

import (
	"context"
	"net/http"
	"testing"

	"github.com/hitnarkhede/learnsphere-api/internal/service"
)

func TestStatsOverview_OK(t *testing.T) {
	e := newEcho()
	h := NewStatsHandler(&mockStatsService{})

	c, rec := newCtx(e, http.MethodGet, "/stats/overview", "")
	if err := h.Overview(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestStatsOverview_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewStatsHandler(&mockStatsService{
		overviewFn: func(_ context.Context) (*service.OverviewResponse, error) {
			return nil, errMock
		},
	})

	c, rec := newCtx(e, http.MethodGet, "/stats/overview", "")
	_ = h.Overview(c)

	assertStatus(t, rec, http.StatusInternalServerError)
}

func TestStatsRevenue_OK(t *testing.T) {
	e := newEcho()
	h := NewStatsHandler(&mockStatsService{})

	c, rec := newCtx(e, http.MethodGet, "/stats/revenue", "")
	if err := h.Revenue(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestStatsRevenue_DefaultPeriod(t *testing.T) {
	e := newEcho()
	var capturedPeriod string
	h := NewStatsHandler(&mockStatsService{
		revenueByPeriodFn: func(_ context.Context, period string) ([]service.PeriodPoint, error) {
			capturedPeriod = period
			return []service.PeriodPoint{}, nil
		},
	})

	c, rec := newCtx(e, http.MethodGet, "/stats/revenue", "")
	if err := h.Revenue(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	if capturedPeriod != "daily" {
		t.Errorf("expected default period=daily, got %s", capturedPeriod)
	}
}

func TestStatsRevenue_CustomPeriod(t *testing.T) {
	e := newEcho()
	var capturedPeriod string
	h := NewStatsHandler(&mockStatsService{
		revenueByPeriodFn: func(_ context.Context, period string) ([]service.PeriodPoint, error) {
			capturedPeriod = period
			return []service.PeriodPoint{}, nil
		},
	})

	req := newEcho()
	_ = req
	c, rec := newCtx(e, http.MethodGet, "/stats/revenue?period=weekly", "")
	if err := h.Revenue(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	if capturedPeriod != "weekly" {
		t.Errorf("expected period=weekly, got %s", capturedPeriod)
	}
}

func TestStatsRevenue_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewStatsHandler(&mockStatsService{
		revenueByPeriodFn: func(_ context.Context, _ string) ([]service.PeriodPoint, error) {
			return nil, errMock
		},
	})

	c, rec := newCtx(e, http.MethodGet, "/stats/revenue", "")
	_ = h.Revenue(c)

	assertStatus(t, rec, http.StatusInternalServerError)
}

func TestStatsActiveUsers_OK(t *testing.T) {
	e := newEcho()
	h := NewStatsHandler(&mockStatsService{})

	c, rec := newCtx(e, http.MethodGet, "/stats/active-users", "")
	if err := h.ActiveUsers(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestStatsActiveUsers_DefaultPeriod(t *testing.T) {
	e := newEcho()
	var capturedPeriod string
	h := NewStatsHandler(&mockStatsService{
		activeUsersByPeriodFn: func(_ context.Context, period string) ([]service.PeriodPoint, error) {
			capturedPeriod = period
			return []service.PeriodPoint{}, nil
		},
	})

	c, rec := newCtx(e, http.MethodGet, "/stats/active-users", "")
	if err := h.ActiveUsers(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	if capturedPeriod != "daily" {
		t.Errorf("expected default period=daily, got %s", capturedPeriod)
	}
}

func TestStatsActiveUsers_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewStatsHandler(&mockStatsService{
		activeUsersByPeriodFn: func(_ context.Context, _ string) ([]service.PeriodPoint, error) {
			return nil, errMock
		},
	})

	c, rec := newCtx(e, http.MethodGet, "/stats/active-users", "")
	_ = h.ActiveUsers(c)

	assertStatus(t, rec, http.StatusInternalServerError)
}
