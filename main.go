package main

import (
	"context"

	"github.com/labstack/echo"

	appRoutes "golang-gu-ask-real-api/routes"
	"golang-gu-ask-real-api/services/mongodb"
)

func main() {
	ctx := context.TODO()
	e := echo.New()

	// Root level middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	// e.Use(middleware.CORS())

	db := mongodb.NewConnectionMongo(ctx)
	defer db.MongoConnect().Disconnect(ctx)

	// Route collection
	appRoutes.RouteCollect(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
