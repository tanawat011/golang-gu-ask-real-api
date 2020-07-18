package controllers

import (
	"golang-gu-ask-real-api/features/crud_with_mongo_mgm/models"
	"net/http"

	"github.com/Kamva/mgm"
	"github.com/labstack/echo"
)

// GetOne : Get one record
func GetOne(c echo.Context) error {
	id := c.QueryParam("id")

	receiver := &models.User{}
	coll := mgm.Coll(receiver)
	_ = coll.FindByID(id, receiver)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Found one",
		"results": receiver,
	})
}
