package controllers

import (
	"golang-gu-ask-real-api/features/crud_with_mongo_driver/models"
	"golang-gu-ask-real-api/utils"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

// Update : Filter and update a found record
func (t *Tools) Update(c echo.Context) error {
	b := new(models.UserUpdate)
	if err := c.Bind(b); err != nil {
		return err
	}

	var filter bson.D
	utils.BsonD(b.Filter, &filter)

	var updater bson.D
	utils.BsonD(b.Update, &updater)
	update := bson.D{
		{"$set", updater},
	}

	if err := t.db.UpdateOne("users", filter, update); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Updated",
	})
}
