package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/marufmax/techtrends/api/pkg/config"
	"net/http"
	"time"
)

type myEcho struct {
	*echo.Echo
}

func New() *myEcho {
	e := myEcho{echo.New()}
	e.defaultMiddlewares()
	// configure default middlewares

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	return &e
}

func (e *myEcho) defaultMiddlewares() *myEcho {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{"GET"},
		AllowOrigins: config.Env.AllowOrigins,
	}))

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	// Rate limiter
	e.Use(middleware.RateLimiterWithConfig(config.RateLimitConfig()))

	// Timeout
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 30 * time.Second,
	}))

	return e
}
