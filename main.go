package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	customMiddleware "golang-gu-ask-real-api/middleware"
)

// ExampleDataBody ...
type ExampleDataBody struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func main() {
	e := echo.New()

	// Root level middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BasicAuth(customMiddleware.BasicAuth))

	e.GET("/:id", ExampleGet)

	e.Logger.Fatal(e.Start(":8080"))
}

// ExampleGet ...
func ExampleGet(c echo.Context) error {
	id := c.Param("id")
	name := c.QueryParam("name")
	return c.String(http.StatusOK, "Example Get with id: "+id+" and name: "+name)
}
