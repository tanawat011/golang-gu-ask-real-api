package crud

import (
	"net/http"

	"github.com/labstack/echo"
)

// Update ...
func Update(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Updated",
	})
}
