// Package main is the entry point for the LearnSphere API server.
//
//	@title			LearnSphere API
//	@version		1.0
//	@description	REST API for the LearnSphere e-learning platform
//	@termsOfService	http://swagger.io/terms/
//
//	@contact.name	LearnSphere Support
//	@contact.email	support@learnsphere.io
//
//	@license.name	MIT
//
//	@host		localhost:8080
//	@BasePath	/api/v1
//
//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Enter: Bearer <JWT token>
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/hitnarkhede/learnsphere-api/internal/config"
	"github.com/hitnarkhede/learnsphere-api/internal/db"
	"github.com/hitnarkhede/learnsphere-api/internal/handler"
	"github.com/hitnarkhede/learnsphere-api/internal/service"
	"github.com/hitnarkhede/learnsphere-api/pkg/validator"

	_ "github.com/hitnarkhede/learnsphere-api/docs"
)

func main() {
	cfg := config.Load()

	pool, err := db.NewPool(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "db connect: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	e := echo.New()
	e.HideBanner = true
	e.Validator = validator.New()

	e.Use(echomiddleware.RequestID())
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.Use(echomiddleware.CORS())
	e.Use(echomiddleware.Secure())

	registerRoutes(e, cfg, pool)

	go func() {
		if err := e.Start(":" + cfg.AppPort); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func registerRoutes(e *echo.Echo, cfg *config.Config, pool *pgxpool.Pool) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	health := handler.NewHealthHandler()
	e.GET("/health", health.Check)

	v1 := e.Group("/api/v1")

	handler.NewAuthHandler(service.NewAuthService(pool, cfg)).RegisterRoutes(v1, cfg)
	handler.NewCourseHandler(service.NewCourseService(pool)).RegisterRoutes(v1, cfg)
	handler.NewEnrollmentHandler(service.NewEnrollmentService(pool)).RegisterRoutes(v1, cfg)
	handler.NewUserHandler(service.NewUserService(pool)).RegisterRoutes(v1, cfg)
	handler.NewTransactionHandler(service.NewTransactionService(pool)).RegisterRoutes(v1, cfg)
	handler.NewWishlistHandler(service.NewWishlistService(pool)).RegisterRoutes(v1, cfg)
	handler.NewChatHandler(service.NewChatService(pool)).RegisterRoutes(v1, cfg)
	handler.NewStatsHandler(service.NewStatsService(pool)).RegisterRoutes(v1, cfg)
	handler.NewFeatureHandler(service.NewFeatureService(pool)).RegisterRoutes(v1, cfg)
	handler.NewLiveSessionHandler(service.NewLiveSessionService(pool)).RegisterRoutes(v1, cfg)
	handler.NewLocaleHandler(service.NewLocaleService(pool)).RegisterRoutes(v1)
}
