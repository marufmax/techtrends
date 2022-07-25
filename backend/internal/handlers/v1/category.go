package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/marufmax/techtrends/api/models"
	"github.com/marufmax/techtrends/api/pkg/postgres"
	"net/http"
)

func GetAllCategories(c echo.Context) error {
	var categories []models.Category

	postgres.Connection().Unscoped().Find(&categories)
	return c.JSON(http.StatusOK, categories)
}
