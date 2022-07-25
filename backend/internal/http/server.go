package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/marufmax/techtrends/api/config"
	"github.com/marufmax/techtrends/api/internal/middlewares"
	"strings"
	"time"
)

type Server struct {
	*echo.Echo
}

func NewApiServer() *Server {
	e := Server{echo.New()}
	// configure default middlewares
	e.defaultMiddlewares()

	// configure app server
	e.Server.ReadTimeout = config.App.ReadTimeout
	e.Server.WriteTimeout = config.App.WriteTimeout

	return &e
}

func (e *Server) defaultMiddlewares() *Server {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{"GET"},
		AllowOrigins: config.Env.AllowOrigins,
	}))

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, config.App.ApiDocPath)
		},
	}))

	// Rate limiter
	e.Use(middleware.RateLimiterWithConfig(middlewares.RateLimitConfig()))

	// Timeout
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 30 * time.Second,
	}))

	return e
}
