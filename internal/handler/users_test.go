package handler

import (
	"context"
	"net/http"
	"testing"

	"github.com/hitnarkhede/learnsphere-api/internal/service"
)

func TestUserList_OK(t *testing.T) {
	e := newEcho()
	h := NewUserHandler(&mockUserService{})

	c, rec := newCtx(e, http.MethodGet, "/users", "")
	if err := h.List(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestUserList_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewUserHandler(&mockUserService{
		listFn: func(_ context.Context) ([]service.UserResponse, error) {
			return nil, errMock
		},
	})

	c, rec := newCtx(e, http.MethodGet, "/users", "")
	_ = h.List(c)

	assertStatus(t, rec, http.StatusInternalServerError)
}

func TestUserDelete_OK(t *testing.T) {
	e := newEcho()
	h := NewUserHandler(&mockUserService{})

	c, rec := newCtx(e, http.MethodDelete, "/users/u1", "")
	c.SetParamNames("id")
	c.SetParamValues("u1")
	if err := h.Delete(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusNoContent)
}

func TestUserDelete_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewUserHandler(&mockUserService{
		deleteFn: func(_ context.Context, _ string) error { return errMock },
	})

	c, rec := newCtx(e, http.MethodDelete, "/users/u1", "")
	c.SetParamNames("id")
	c.SetParamValues("u1")
	_ = h.Delete(c)

	assertStatus(t, rec, http.StatusInternalServerError)
}
