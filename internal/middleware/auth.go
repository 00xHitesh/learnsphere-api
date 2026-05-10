package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/user/learnsphere-api/internal/config"
	"github.com/user/learnsphere-api/pkg/jwt"
	"github.com/user/learnsphere-api/pkg/response"
)

func JWT(cfg *config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			header := c.Request().Header.Get("Authorization")
			if !strings.HasPrefix(header, "Bearer ") {
				return response.Unauthorized(c, "missing or malformed token")
			}

			tokenStr := strings.TrimPrefix(header, "Bearer ")
			claims, err := jwt.ValidatToken(tokenStr, cfg.JWTAccessSecret)
			if err != nil {
				return response.Unauthorized(c, "invalid or expired token")
			}

			c.Set("user_id", claims.UserID)
			c.Set("email", claims.Email)
			c.Set("role", claims.Role)

			return next(c)
		}
	}
}

func RequireRole(roles ...string) echo.MiddlewareFunc {
	allowed := make(map[string]struct{}, len(roles))
	for _, r := range roles {
		allowed[r] = struct{}{}
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			role, _ := c.Get("role").(string)
			if _, ok := allowed[role]; !ok {
				return response.Forbidden(c, "insufficient permissions")
			}
			return next(c)
		}
	}
}
