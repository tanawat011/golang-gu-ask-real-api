package main

import (
	"context"

	"github.com/labstack/echo"

	appRoutes "golang-gu-ask-real-api/routes"
	"golang-gu-ask-real-api/services/mongodb"
	mongodbMGM "golang-gu-ask-real-api/services/mongodb_mgm"
)

const (
	uri    = "mongodb://localhost:27017"
	dbName = "example_crud"
)

func main() {
	ctx := context.TODO()
	e := echo.New()

	// Root level middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	// e.Use(middleware.CORS())

	// Mongo driver official
	db := mongodb.NewConnectionMongo(ctx)
	defer db.MongoConnect().Disconnect(ctx)

	// MGM odm base on mongo driver official
	mongodbMGM.NewConnectionMongo()

	// Route collection
	appRoutes.RouteCollect(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
