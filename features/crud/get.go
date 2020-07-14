package crud

import (
	"net/http"

	"github.com/labstack/echo"
)

// Get ...
func Get(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Get done",
	})
}
