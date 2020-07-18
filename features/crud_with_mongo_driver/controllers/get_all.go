package controllers

import (
	"golang-gu-ask-real-api/features/crud_with_mongo_driver/models"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAll ...
func (t *Tools) GetAll(c echo.Context) error {
	filter := bson.D{}

	receiver := make([]map[string]interface{}, 0)

	if err := t.db.Find("users", filter, &receiver); err != nil {
		return err
	}

	response := make(models.UserReceiverList, 0)
	for _, v := range receiver {
		response = append(response, &models.UserReceiver{
			ID: v["_id"].(primitive.ObjectID),
		})
	}
	// mapstructure.Decode(receiver, &response)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Found",
		"data":    response,
	})
}
