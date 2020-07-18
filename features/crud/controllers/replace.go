package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

// Replace ...
func Replace(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Replaced",
	})
}
