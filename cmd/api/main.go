package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/user/learnsphere-api/internal/config"
	"github.com/user/learnsphere-api/internal/db"
	"github.com/user/learnsphere-api/internal/handler"
	"github.com/user/learnsphere-api/pkg/validator"
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

	registerRoutes(e, cfg)

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

func registerRoutes(e *echo.Echo, cfg *config.Config) {
	health := handler.NewHealthHandler()

	e.GET("/health", health.Check)

	v1 := e.Group("/api/v1")
	_ = v1
	// Register domain route groups here, e.g.:
	// auth.RegisterRoutes(v1, cfg, pool)
	// users.RegisterRoutes(v1, cfg, pool)
	// courses.RegisterRoutes(v1, cfg, pool)
}
