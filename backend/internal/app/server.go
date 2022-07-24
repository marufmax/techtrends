package app

import (
	"github.com/labstack/echo/v4"
	"github.com/marufmax/techtrends/api/pkg/bugsnug"
	"github.com/marufmax/techtrends/api/pkg/config"
	"net/http"
)

func Run() {
	bugsnug.Configure()
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":" + config.Env.AppPort))
}
