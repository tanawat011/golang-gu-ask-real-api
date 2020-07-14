package crud

import (
	"net/http"

	"github.com/labstack/echo"
)

// Delete ...
func Delete(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Deleted",
	})
}
