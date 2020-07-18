package controllers

import (
	"golang-gu-ask-real-api/features/crud_with_mongo_mgm/models"
	"net/http"

	"github.com/Kamva/mgm"
	"github.com/Kamva/mgm/operator"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

// GetByFilter : Get record by filter
func GetByFilter(c echo.Context) error {
	name := "name"
	val := "Tanawat"
	receiver := []models.User{}
	coll := mgm.Coll(&models.User{})
	_ = coll.SimpleFind(&receiver, bson.M{name: bson.M{operator.Gte: val}})

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Found one",
		"results": receiver,
	})
}
