package handler

import (
	"context"
	"net/http"
	"testing"

	"github.com/hitnarkhede/learnsphere-api/internal/service"
)

func TestFeatureList_OK(t *testing.T) {
	e := newEcho()
	h := NewFeatureHandler(&mockFeatureService{})

	c, rec := newCtx(e, http.MethodGet, "/features", "")
	if err := h.List(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestFeatureList_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewFeatureHandler(&mockFeatureService{
		listFn: func(_ context.Context) ([]service.FeatureFlagResponse, error) {
			return nil, errMock
		},
	})

	c, rec := newCtx(e, http.MethodGet, "/features", "")
	_ = h.List(c)

	assertStatus(t, rec, http.StatusInternalServerError)
}

func TestFeatureToggle_OK(t *testing.T) {
	e := newEcho()
	h := NewFeatureHandler(&mockFeatureService{})

	c, rec := newCtx(e, http.MethodPatch, "/features/dark-mode/toggle", "")
	c.SetParamNames("name")
	c.SetParamValues("dark-mode")
	if err := h.Toggle(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestFeatureToggle_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewFeatureHandler(&mockFeatureService{
		toggleFn: func(_ context.Context, _ string) ([]service.FeatureFlagResponse, error) {
			return nil, errMock
		},
	})

	c, rec := newCtx(e, http.MethodPatch, "/features/dark-mode/toggle", "")
	c.SetParamNames("name")
	c.SetParamValues("dark-mode")
	_ = h.Toggle(c)

	assertStatus(t, rec, http.StatusInternalServerError)
}
