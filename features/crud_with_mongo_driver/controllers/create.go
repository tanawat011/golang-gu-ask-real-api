package controllers

import (
	"golang-gu-ask-real-api/features/crud_with_mongo_driver/models"
	"net/http"

	"github.com/labstack/echo"
)

// Create : Create a new data
func (t *Tools) Create(c echo.Context) error {
	b := new(models.UserCreate)
	if err := c.Bind(b); err != nil {
		return err
	}

	if err := t.db.InsertOne("users", b); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Created",
	})
}
