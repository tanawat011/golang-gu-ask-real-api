package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	appRoutes "golang-gu-ask-real-api/routes"
)

func main() {
	e := echo.New()

	// Root level middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Route collection
	appRoutes.RouteCollect(e)

	e.Logger.Fatal(e.Start(":8080"))
}
