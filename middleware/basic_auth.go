package middleware

import (
	"github.com/labstack/echo"
)

// BasicAuth ...
func BasicAuth(username, password string, c echo.Context) (bool, error) {
	if username == "tanawat" && password == "198777" {
		return true, nil
	}
	return false, nil
}
