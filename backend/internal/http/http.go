package http

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/marufmax/techtrends/api/cmd/docs"
	"github.com/marufmax/techtrends/api/config"
	v1 "github.com/marufmax/techtrends/api/internal/handlers/v1"

	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

func Run() {
	e := NewApiServer()
	docs.SwaggerInfo.BasePath = "/"
	// Default root path
	e.GET("/", func(c echo.Context) error {
		res := `{"hi": "nice to see you!"}`
		var resMap map[string]interface{}
		json.Unmarshal([]byte(res), &resMap)
		return c.JSON(http.StatusOK, resMap)
	})

	// Swagger route
	e.GET(config.App.ApiDocPath, echoSwagger.WrapHandler)
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hi")
	})

	// Application routes
	e.GET("/categories", v1.GetAllCategories)

	e.Logger.Fatal(e.Start(":" + config.Env.AppPort))
}
