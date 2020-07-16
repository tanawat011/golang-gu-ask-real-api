package crud

import (
	"fmt"
	"golang-gu-ask-real-api/features/crud/entities"
	"golang-gu-ask-real-api/services/mongodb"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

// GetAll ...
func GetAll(c echo.Context) error {
	filter := bson.D{}

	// var receiver entities.UserReceiverList
	receiver := make(entities.UserReceiverList, 0)

	if err := mongodb.Find("users", filter, &receiver); err != nil {
		return err
	}
	for _, v := range receiver {
		fmt.Println(v)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Found",
		"results": receiver,
	})
}
