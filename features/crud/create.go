package crud

import (
	"net/http"

	"github.com/labstack/echo"
)

// Create ...
func Create(c echo.Context) error {
	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Created",
	})
}
