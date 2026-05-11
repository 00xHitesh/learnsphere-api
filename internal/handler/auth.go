package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/hitnarkhede/learnsphere-api/internal/config"
	"github.com/hitnarkhede/learnsphere-api/internal/middleware"
	"github.com/hitnarkhede/learnsphere-api/pkg/response"
)

type AuthHandler struct {
	svc authService
}

func NewAuthHandler(svc authService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

func (h *AuthHandler) RegisterRoutes(v1 *echo.Group, cfg *config.Config) {
	g := v1.Group("/auth")
	g.POST("/login", h.Login)
	g.POST("/register", h.Register)
	g.POST("/reset-request", h.RequestReset)
	g.POST("/reset-password", h.ResetPassword)
	g.POST("/refresh", h.Refresh)

	secure := g.Group("", middleware.JWT(cfg))
	secure.GET("/me", h.Me)
	secure.PUT("/interests", h.UpdateInterests)
	secure.PUT("/locale", h.UpdateLocale)
}

// Request types

type LoginRequest struct {
	Email    string `json:"email" example:"demo@lumen.io"`
	Password string `json:"password" example:"secret123"`
}

type RegisterRequest struct {
	Name     string `json:"name" example:"Jane Doe"`
	Email    string `json:"email" example:"jane@example.com"`
	Password string `json:"password" example:"secret123"`
}

type UpdateInterestsRequest struct {
	Interests []string `json:"interests" example:"Design,Programming"`
}

type UpdateLocaleRequest struct {
	RegionCode   string `json:"regionCode" example:"US"`
	CurrencyCode string `json:"currencyCode" example:"USD"`
}

type ResetRequestBody struct {
	Email string `json:"email" example:"demo@lumen.io"`
}

type ResetPasswordRequest struct {
	Token    string `json:"token" example:"abc123def456"`
	Password string `json:"password" example:"newpassword123"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

// Refresh godoc
//
//	@Summary		Refresh tokens
//	@Description	Exchange a valid refresh token for a new access + refresh token pair
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		RefreshTokenRequest	true	"Refresh token"
//	@Success		200		{object}	response.Response{data=service.TokenPair}
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		401		{object}	response.ErrorResponse
//	@Router			/auth/refresh [post]
func (h *AuthHandler) Refresh(c echo.Context) error {
	var body struct {
		RefreshToken string `json:"refreshToken" validate:"required"`
	}
	if err := bindAndValidate(c, &body); err != nil {
		return err
	}
	tokens, err := h.svc.RefreshToken(c.Request().Context(), body.RefreshToken)
	if err != nil {
		return response.Unauthorized(c, err.Error())
	}
	return response.OK(c, tokens)
}

// Login godoc
//
//	@Summary		Login
//	@Description	Authenticate with email and password, returns user + JWT token pair
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		LoginRequest	true	"Login credentials"
//	@Success		200		{object}	response.Response{data=object{user=service.UserResponse,tokens=service.TokenPair}}
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		401		{object}	response.ErrorResponse
//	@Router			/auth/login [post]
func (h *AuthHandler) Login(c echo.Context) error {
	var body struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
	if err := bindAndValidate(c, &body); err != nil {
		return err
	}
	user, tokens, err := h.svc.Login(c.Request().Context(), body.Email, body.Password)
	if err != nil {
		return response.Unauthorized(c, err.Error())
	}
	return response.OK(c, map[string]interface{}{"user": user, "tokens": tokens})
}

// Register godoc
//
//	@Summary		Register
//	@Description	Create a new user account, returns user + JWT token pair
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		RegisterRequest	true	"Registration details"
//	@Success		201		{object}	response.Response{data=object{user=service.UserResponse,tokens=service.TokenPair}}
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		409		{object}	response.ErrorResponse	"Email already in use"
//	@Router			/auth/register [post]
func (h *AuthHandler) Register(c echo.Context) error {
	var body struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6"`
	}
	if err := bindAndValidate(c, &body); err != nil {
		return err
	}
	user, tokens, err := h.svc.Register(c.Request().Context(), body.Name, body.Email, body.Password)
	if err != nil {
		return response.Conflict(c, err.Error())
	}
	return response.Created(c, map[string]interface{}{"user": user, "tokens": tokens})
}

// Me godoc
//
//	@Summary		Get current user
//	@Description	Returns the authenticated user's profile including interests
//	@Tags			Auth
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{object}	response.Response{data=service.UserResponse}
//	@Failure		401	{object}	response.ErrorResponse
//	@Failure		404	{object}	response.ErrorResponse
//	@Router			/auth/me [get]
func (h *AuthHandler) Me(c echo.Context) error {
	userID := c.Get("user_id").(string)
	user, err := h.svc.GetUser(c.Request().Context(), userID)
	if err != nil {
		return response.NotFound(c, "user not found")
	}
	return response.OK(c, user)
}

// UpdateInterests godoc
//
//	@Summary		Update interests
//	@Description	Replace the authenticated user's interest categories
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			body	body		UpdateInterestsRequest	true	"Interest list"
//	@Success		200		{object}	response.Response{data=service.UserResponse}
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		401		{object}	response.ErrorResponse
//	@Router			/auth/interests [put]
func (h *AuthHandler) UpdateInterests(c echo.Context) error {
	var body struct {
		Interests []string `json:"interests" validate:"required"`
	}
	if err := bindAndValidate(c, &body); err != nil {
		return err
	}
	userID := c.Get("user_id").(string)
	user, err := h.svc.UpdateInterests(c.Request().Context(), userID, body.Interests)
	if err != nil {
		return response.InternalError(c, "")
	}
	return response.OK(c, user)
}

// UpdateLocale godoc
//
//	@Summary		Update locale
//	@Description	Update the authenticated user's region and currency preference
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			body	body		UpdateLocaleRequest	true	"Locale settings"
//	@Success		200		{object}	response.Response{data=service.UserResponse}
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		401		{object}	response.ErrorResponse
//	@Router			/auth/locale [put]
func (h *AuthHandler) UpdateLocale(c echo.Context) error {
	var body struct {
		RegionCode   string `json:"regionCode" validate:"required"`
		CurrencyCode string `json:"currencyCode" validate:"required"`
	}
	if err := bindAndValidate(c, &body); err != nil {
		return err
	}
	userID := c.Get("user_id").(string)
	user, err := h.svc.UpdateLocale(c.Request().Context(), userID, body.RegionCode, body.CurrencyCode)
	if err != nil {
		return response.InternalError(c, "")
	}
	return response.OK(c, user)
}

// RequestReset godoc
//
//	@Summary		Request password reset
//	@Description	Send a password-reset token for the given email
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		ResetRequestBody	true	"Email address"
//	@Success		200		{object}	response.Response{data=object{token=string}}
//	@Failure		400		{object}	response.ErrorResponse
//	@Router			/auth/reset-request [post]
func (h *AuthHandler) RequestReset(c echo.Context) error {
	var body struct {
		Email string `json:"email" validate:"required,email"`
	}
	if err := bindAndValidate(c, &body); err != nil {
		return err
	}
	token, err := h.svc.RequestReset(c.Request().Context(), body.Email)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}
	return response.OK(c, map[string]string{"token": token})
}

// ResetPassword godoc
//
//	@Summary		Reset password
//	@Description	Set a new password using a valid reset token
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body	ResetPasswordRequest	true	"Token and new password"
//	@Success		204		"No content"
//	@Failure		400		{object}	response.ErrorResponse
//	@Router			/auth/reset-password [post]
func (h *AuthHandler) ResetPassword(c echo.Context) error {
	var body struct {
		Token    string `json:"token" validate:"required"`
		Password string `json:"password" validate:"required,min=6"`
	}
	if err := bindAndValidate(c, &body); err != nil {
		return err
	}
	if err := h.svc.ResetPassword(c.Request().Context(), body.Token, body.Password); err != nil {
		return response.BadRequest(c, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
