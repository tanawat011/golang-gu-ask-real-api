package crud

import (
	"fmt"
	"golang-gu-ask-real-api/features/crud/entities"
	"golang-gu-ask-real-api/services/mongodb"
	"golang-gu-ask-real-api/utils"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

// Update : Filter and update a found record
func Update(c echo.Context) error {
	b := new(entities.UserUpdate)
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
	fmt.Println(filter)
	fmt.Println(updater)

	if err := mongodb.UpdateOne("users", filter, update); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Updated",
	})
}
