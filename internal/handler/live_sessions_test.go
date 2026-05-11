package handler

import (
	"context"
	"net/http"
	"testing"

	"github.com/hitnarkhede/learnsphere-api/internal/service"
)

func TestLiveSessionList_OK(t *testing.T) {
	e := newEcho()
	h := NewLiveSessionHandler(&mockLiveSessionService{})

	c, rec := newCtx(e, http.MethodGet, "/live-sessions", "")
	if err := h.List(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestLiveSessionList_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewLiveSessionHandler(&mockLiveSessionService{
		listPublishedFn: func(_ context.Context) ([]service.LiveSessionResponse, error) {
			return nil, errMock
		},
	})

	c, rec := newCtx(e, http.MethodGet, "/live-sessions", "")
	_ = h.List(c)

	assertStatus(t, rec, http.StatusInternalServerError)
}

func TestLiveSessionListAll_OK(t *testing.T) {
	e := newEcho()
	h := NewLiveSessionHandler(&mockLiveSessionService{})

	c, rec := newCtx(e, http.MethodGet, "/live-sessions/all", "")
	if err := h.ListAll(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestLiveSessionListAll_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewLiveSessionHandler(&mockLiveSessionService{
		listAllFn: func(_ context.Context) ([]service.LiveSessionResponse, error) {
			return nil, errMock
		},
	})

	c, rec := newCtx(e, http.MethodGet, "/live-sessions/all", "")
	_ = h.ListAll(c)

	assertStatus(t, rec, http.StatusInternalServerError)
}

func TestLiveSessionUpsert_OK(t *testing.T) {
	e := newEcho()
	h := NewLiveSessionHandler(&mockLiveSessionService{})

	body := `{
		"title":"TS Deep Dive","instructor":"Alice","scheduledAt":"2025-09-01T10:00:00Z",
		"durationMinutes":90,"joinUrl":"https://meet.example.com/abc","maxAttendees":100,
		"status":"published"
	}`
	c, rec := newCtx(e, http.MethodPut, "/live-sessions", body)
	if err := h.Upsert(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestLiveSessionUpsert_ValidationError(t *testing.T) {
	e := newEcho()
	h := NewLiveSessionHandler(&mockLiveSessionService{})

	// missing required fields
	c, rec := newCtx(e, http.MethodPut, "/live-sessions", `{"title":"Only Title"}`)
	_ = h.Upsert(c)

	assertStatus(t, rec, http.StatusBadRequest)
}

func TestLiveSessionUpsert_InvalidScheduledAt(t *testing.T) {
	e := newEcho()
	h := NewLiveSessionHandler(&mockLiveSessionService{})

	body := `{
		"title":"TS Deep Dive","instructor":"Alice","scheduledAt":"not-a-date",
		"durationMinutes":90,"joinUrl":"https://meet.example.com/abc","maxAttendees":100,
		"status":"published"
	}`
	c, rec := newCtx(e, http.MethodPut, "/live-sessions", body)
	_ = h.Upsert(c)

	assertStatus(t, rec, http.StatusBadRequest)
}

func TestLiveSessionUpsert_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewLiveSessionHandler(&mockLiveSessionService{
		upsertFn: func(_ context.Context, _ service.UpsertLiveSessionParams) (*service.LiveSessionResponse, error) {
			return nil, errMock
		},
	})

	body := `{
		"title":"TS Deep Dive","instructor":"Alice","scheduledAt":"2025-09-01T10:00:00Z",
		"durationMinutes":90,"joinUrl":"https://meet.example.com/abc","maxAttendees":100,
		"status":"published"
	}`
	c, rec := newCtx(e, http.MethodPut, "/live-sessions", body)
	_ = h.Upsert(c)

	assertStatus(t, rec, http.StatusInternalServerError)
}

func TestLiveSessionDelete_OK(t *testing.T) {
	e := newEcho()
	h := NewLiveSessionHandler(&mockLiveSessionService{})

	c, rec := newCtx(e, http.MethodDelete, "/live-sessions/s1", "")
	c.SetParamNames("id")
	c.SetParamValues("s1")
	if err := h.Delete(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusNoContent)
}

func TestLiveSessionDelete_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewLiveSessionHandler(&mockLiveSessionService{
		deleteFn: func(_ context.Context, _ string) error { return errMock },
	})

	c, rec := newCtx(e, http.MethodDelete, "/live-sessions/s1", "")
	c.SetParamNames("id")
	c.SetParamValues("s1")
	_ = h.Delete(c)

	assertStatus(t, rec, http.StatusInternalServerError)
}

func TestLiveSessionSetStatus_OK(t *testing.T) {
	e := newEcho()
	h := NewLiveSessionHandler(&mockLiveSessionService{})

	c, rec := newCtx(e, http.MethodPatch, "/live-sessions/s1/status", `{"status":"published"}`)
	c.SetParamNames("id")
	c.SetParamValues("s1")
	if err := h.SetStatus(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestLiveSessionSetStatus_ValidationError(t *testing.T) {
	e := newEcho()
	h := NewLiveSessionHandler(&mockLiveSessionService{})

	c, rec := newCtx(e, http.MethodPatch, "/live-sessions/s1/status", `{}`)
	c.SetParamNames("id")
	c.SetParamValues("s1")
	_ = h.SetStatus(c)

	assertStatus(t, rec, http.StatusBadRequest)
}

func TestLiveSessionSetStatus_NotFound(t *testing.T) {
	e := newEcho()
	h := NewLiveSessionHandler(&mockLiveSessionService{
		setStatusFn: func(_ context.Context, _, _ string) (*service.LiveSessionResponse, error) {
			return nil, nil
		},
	})

	c, rec := newCtx(e, http.MethodPatch, "/live-sessions/s1/status", `{"status":"published"}`)
	c.SetParamNames("id")
	c.SetParamValues("s1")
	_ = h.SetStatus(c)

	assertStatus(t, rec, http.StatusNotFound)
}

func TestLiveSessionSetStatus_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewLiveSessionHandler(&mockLiveSessionService{
		setStatusFn: func(_ context.Context, _, _ string) (*service.LiveSessionResponse, error) {
			return nil, errMock
		},
	})

	c, rec := newCtx(e, http.MethodPatch, "/live-sessions/s1/status", `{"status":"published"}`)
	c.SetParamNames("id")
	c.SetParamValues("s1")
	_ = h.SetStatus(c)

	assertStatus(t, rec, http.StatusInternalServerError)
}
