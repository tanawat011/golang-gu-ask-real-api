package crud

import (
	"golang-gu-ask-real-api/features/crud/entities"
	"golang-gu-ask-real-api/services/mongodb"
	"net/http"

	"github.com/labstack/echo"
)

// Create : Create a new data
func Create(c echo.Context) error {
	b := new(entities.UserCreate)
	if err := c.Bind(b); err != nil {
		return err
	}

	if err := mongodb.InsertOne("users", b); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Created",
	})
}
