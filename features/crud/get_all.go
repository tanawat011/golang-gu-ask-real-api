package crud

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetAll ...
func GetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Get all done",
	})
}
