package handler

import (
	"net/http"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	e := newEcho()
	h := NewHealthHandler()

	c, rec := newCtx(e, http.MethodGet, "/health", "")
	if err := h.Check(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}
