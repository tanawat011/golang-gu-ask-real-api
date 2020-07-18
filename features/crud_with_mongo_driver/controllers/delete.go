package controllers

import (
	"golang-gu-ask-real-api/features/crud_with_mongo_driver/models"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Delete ...
func (t *Tools) Delete(c echo.Context) error {
	b := new(models.UserFind)
	if err := c.Bind(b); err != nil {
		return err
	}

	objID, _ := primitive.ObjectIDFromHex(b.ID)
	filter := bson.D{{"_id", objID}}

	if err := t.db.Delete("users", filter); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Deleted",
	})
}
