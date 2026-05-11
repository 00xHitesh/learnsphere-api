package handler

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	customvalidator "github.com/hitnarkhede/learnsphere-api/pkg/validator"
)

func newEcho() *echo.Echo {
	e := echo.New()
	e.Validator = customvalidator.New()
	return e
}

func newCtx(e *echo.Echo, method, path, body string) (echo.Context, *httptest.ResponseRecorder) {
	var reqBody io.Reader
	if body != "" {
		reqBody = strings.NewReader(body)
	}
	req := httptest.NewRequest(method, path, reqBody)
	if body != "" {
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	}
	rec := httptest.NewRecorder()
	return e.NewContext(req, rec), rec
}

func withUser(c echo.Context, userID string) echo.Context {
	c.Set("user_id", userID)
	return c
}

func decodeBody(t *testing.T, rec *httptest.ResponseRecorder) map[string]interface{} {
	t.Helper()
	var out map[string]interface{}
	if err := json.NewDecoder(rec.Body).Decode(&out); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	return out
}

func assertStatus(t *testing.T, rec *httptest.ResponseRecorder, want int) {
	t.Helper()
	if rec.Code != want {
		t.Errorf("status = %d, want %d; body: %s", rec.Code, want, rec.Body)
	}
}

func assertSuccess(t *testing.T, rec *httptest.ResponseRecorder) {
	t.Helper()
	body := decodeBody(t, rec)
	if success, _ := body["success"].(bool); !success {
		t.Errorf("expected success=true, body: %v", body)
	}
}
