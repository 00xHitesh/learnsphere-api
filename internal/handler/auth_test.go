package handler

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/hitnarkhede/learnsphere-api/internal/service"
)

func TestAuthLogin_OK(t *testing.T) {
	e := newEcho()
	h := NewAuthHandler(&mockAuthService{})

	c, rec := newCtx(e, http.MethodPost, "/auth/login", `{"email":"demo@lumen.io","password":"Test@123"}`)
	if err := h.Login(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestAuthLogin_ValidationError(t *testing.T) {
	e := newEcho()
	h := NewAuthHandler(&mockAuthService{})

	c, rec := newCtx(e, http.MethodPost, "/auth/login", `{"email":"not-an-email","password":"pass"}`)
	_ = h.Login(c)

	assertStatus(t, rec, http.StatusBadRequest)
}

func TestAuthLogin_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewAuthHandler(&mockAuthService{
		loginFn: func(_ context.Context, _, _ string) (*service.UserResponse, *service.TokenPair, error) {
			return nil, nil, errors.New("invalid email or password")
		},
	})

	c, rec := newCtx(e, http.MethodPost, "/auth/login", `{"email":"demo@lumen.io","password":"wrong"}`)
	_ = h.Login(c)

	assertStatus(t, rec, http.StatusUnauthorized)
}

func TestAuthRegister_OK(t *testing.T) {
	e := newEcho()
	h := NewAuthHandler(&mockAuthService{})

	c, rec := newCtx(e, http.MethodPost, "/auth/register", `{"name":"Jane","email":"jane@example.com","password":"secret123"}`)
	if err := h.Register(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusCreated)
	assertSuccess(t, rec)
}

func TestAuthRegister_ValidationError(t *testing.T) {
	e := newEcho()
	h := NewAuthHandler(&mockAuthService{})

	// missing password
	c, rec := newCtx(e, http.MethodPost, "/auth/register", `{"name":"Jane","email":"jane@example.com"}`)
	_ = h.Register(c)

	assertStatus(t, rec, http.StatusBadRequest)
}

func TestAuthRegister_Conflict(t *testing.T) {
	e := newEcho()
	h := NewAuthHandler(&mockAuthService{
		registerFn: func(_ context.Context, _, _, _ string) (*service.UserResponse, *service.TokenPair, error) {
			return nil, nil, errors.New("email already in use")
		},
	})

	c, rec := newCtx(e, http.MethodPost, "/auth/register", `{"name":"Jane","email":"jane@example.com","password":"secret123"}`)
	_ = h.Register(c)

	assertStatus(t, rec, http.StatusConflict)
}

func TestAuthMe_OK(t *testing.T) {
	e := newEcho()
	h := NewAuthHandler(&mockAuthService{})

	c, rec := newCtx(e, http.MethodGet, "/auth/me", "")
	withUser(c, "user-123")
	if err := h.Me(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestAuthMe_NotFound(t *testing.T) {
	e := newEcho()
	h := NewAuthHandler(&mockAuthService{
		getUserFn: func(_ context.Context, _ string) (*service.UserResponse, error) {
			return nil, errors.New("not found")
		},
	})

	c, rec := newCtx(e, http.MethodGet, "/auth/me", "")
	withUser(c, "user-123")
	_ = h.Me(c)

	assertStatus(t, rec, http.StatusNotFound)
}

func TestAuthUpdateInterests_OK(t *testing.T) {
	e := newEcho()
	h := NewAuthHandler(&mockAuthService{})

	c, rec := newCtx(e, http.MethodPut, "/auth/interests", `{"interests":["Go","TypeScript"]}`)
	withUser(c, "user-123")
	if err := h.UpdateInterests(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestAuthUpdateInterests_ValidationError(t *testing.T) {
	e := newEcho()
	h := NewAuthHandler(&mockAuthService{})

	// interests field missing
	c, rec := newCtx(e, http.MethodPut, "/auth/interests", `{}`)
	withUser(c, "user-123")
	_ = h.UpdateInterests(c)

	assertStatus(t, rec, http.StatusBadRequest)
}

func TestAuthUpdateInterests_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewAuthHandler(&mockAuthService{
		updateInterestsFn: func(_ context.Context, _ string, _ []string) (*service.UserResponse, error) {
			return nil, errMock
		},
	})

	c, rec := newCtx(e, http.MethodPut, "/auth/interests", `{"interests":["Go"]}`)
	withUser(c, "user-123")
	_ = h.UpdateInterests(c)

	assertStatus(t, rec, http.StatusInternalServerError)
}

func TestAuthUpdateLocale_OK(t *testing.T) {
	e := newEcho()
	h := NewAuthHandler(&mockAuthService{})

	c, rec := newCtx(e, http.MethodPut, "/auth/locale", `{"regionCode":"US","currencyCode":"USD"}`)
	withUser(c, "user-123")
	if err := h.UpdateLocale(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestAuthUpdateLocale_ValidationError(t *testing.T) {
	e := newEcho()
	h := NewAuthHandler(&mockAuthService{})

	c, rec := newCtx(e, http.MethodPut, "/auth/locale", `{"regionCode":"US"}`)
	withUser(c, "user-123")
	_ = h.UpdateLocale(c)

	assertStatus(t, rec, http.StatusBadRequest)
}

func TestAuthRequestReset_OK(t *testing.T) {
	e := newEcho()
	h := NewAuthHandler(&mockAuthService{})

	c, rec := newCtx(e, http.MethodPost, "/auth/reset-request", `{"email":"demo@lumen.io"}`)
	if err := h.RequestReset(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestAuthRequestReset_ValidationError(t *testing.T) {
	e := newEcho()
	h := NewAuthHandler(&mockAuthService{})

	c, rec := newCtx(e, http.MethodPost, "/auth/reset-request", `{"email":"not-an-email"}`)
	_ = h.RequestReset(c)

	assertStatus(t, rec, http.StatusBadRequest)
}

func TestAuthRequestReset_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewAuthHandler(&mockAuthService{
		requestResetFn: func(_ context.Context, _ string) (string, error) {
			return "", errors.New("email not found")
		},
	})

	c, rec := newCtx(e, http.MethodPost, "/auth/reset-request", `{"email":"unknown@example.com"}`)
	_ = h.RequestReset(c)

	assertStatus(t, rec, http.StatusBadRequest)
}

func TestAuthResetPassword_OK(t *testing.T) {
	e := newEcho()
	h := NewAuthHandler(&mockAuthService{})

	c, rec := newCtx(e, http.MethodPost, "/auth/reset-password", `{"token":"abc123","password":"newpass1"}`)
	if err := h.ResetPassword(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusNoContent)
}

func TestAuthResetPassword_ValidationError(t *testing.T) {
	e := newEcho()
	h := NewAuthHandler(&mockAuthService{})

	// password too short (min=6)
	c, rec := newCtx(e, http.MethodPost, "/auth/reset-password", `{"token":"abc","password":"ab"}`)
	_ = h.ResetPassword(c)

	assertStatus(t, rec, http.StatusBadRequest)
}

func TestAuthResetPassword_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewAuthHandler(&mockAuthService{
		resetPasswordFn: func(_ context.Context, _, _ string) error {
			return errors.New("invalid or expired token")
		},
	})

	c, rec := newCtx(e, http.MethodPost, "/auth/reset-password", `{"token":"bad-token","password":"newpass1"}`)
	_ = h.ResetPassword(c)

	assertStatus(t, rec, http.StatusBadRequest)
}
