package crud

import (
	"golang-gu-ask-real-api/features/crud/entities"
	"golang-gu-ask-real-api/services/mongodb"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAll ...
func GetAll(c echo.Context) error {
	b := new(entities.UserFind)
	if err := c.Bind(b); err != nil {
		return err
	}

	objID, _ := primitive.ObjectIDFromHex(b.ID)
	filter := bson.D{{"_id", objID}}

	var receiver entities.UserReceiverList

	if err := mongodb.Find("users", filter, &receiver); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Found",
		"results": receiver,
	})
}
