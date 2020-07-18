package controllers

import (
	"golang-gu-ask-real-api/features/crud_with_mongo_driver/models"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get : Get one record
func (t *Tools) Get(c echo.Context) error {
	b := new(models.UserFind)
	if err := c.Bind(b); err != nil {
		return err
	}

	objID, _ := primitive.ObjectIDFromHex(b.ID)
	filter := bson.D{{"_id", objID}}

	var receiver models.UserReceiver

	if err := t.db.FindOne("users", filter, &receiver); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Found one",
		"results": receiver,
	})
}
