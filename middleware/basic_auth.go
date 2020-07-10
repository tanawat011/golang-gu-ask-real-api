package middleware

import (
	"fmt"

	"github.com/labstack/echo"
)

// BasicAuth ...
func BasicAuth(username, password string, c echo.Context) (bool, error) {
	fmt.Printf("Username: %s, Password: %s", username, password)
	if username == "tanawat" && password == "198777" {
		return true, nil
	}
	return false, nil
}
