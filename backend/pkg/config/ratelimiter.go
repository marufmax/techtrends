package config

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

func RateLimitConfig() middleware.RateLimiterConfig {
	return middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()

			return id, nil
		},
		ErrorHandler: func(ctx echo.Context, err error) error {
			return ctx.JSON(http.StatusForbidden, struct {
				message string
			}{
				message: "Its our fault. sorry!",
			})
		},
		DenyHandler: func(ctx echo.Context, identifier string, err error) error {
			return ctx.JSON(http.StatusTooManyRequests, struct {
				message string
			}{
				message: "You are going too fast! slow down",
			})
		},

		Store: middleware.NewRateLimiterMemoryStoreWithConfig(middleware.RateLimiterMemoryStoreConfig{
			Rate:      10,
			Burst:     30,
			ExpiresIn: 3 * time.Minute,
		}),
	}
}
