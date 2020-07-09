package main

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	// Group level middleware
	g := e.Group("/body")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "tanawat" && password == "198777" {
			return true, nil
		}
		return false, nil
	}))

	// Route level middleware
	track := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			println("request to root route")
			return next(c)
		}
	}

	e.GET("/:id", ExampleGet, track)
	e.POST("/form-urlencoded", ExamplePostForm)
	e.POST("/multipart-form-data", ExamplePostMultiPart)
	g.POST("/json", ExamplePostJSONBody)
	g.POST("/xml", ExamplePostXMLBody)
	g.POST("/form-data", ExamplePostFormDataBody)

	e.Logger.Fatal(e.Start(":8080"))
}

// ExampleGet ...
func ExampleGet(c echo.Context) error {
	id := c.Param("id")
	name := c.QueryParam("name")
	return c.String(http.StatusOK, "Example Get with id: "+id+" and name: "+name)
}

// ExamplePostJSONBody ...
func ExamplePostJSONBody(c echo.Context) error {
	d := new(ExampleDataBody)
	if err := c.Bind(d); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, d)
}

// ExamplePostXMLBody ...
func ExamplePostXMLBody(c echo.Context) error {
	d := new(ExampleDataBody)
	if err := c.Bind(d); err != nil {
		return err
	}
	return c.XML(http.StatusCreated, d)
}

// ExamplePostFormDataBody ...
func ExamplePostFormDataBody(c echo.Context) error {
	d := new(ExampleDataBody)
	if err := c.Bind(d); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, d)
}

// ExamplePostForm ...
func ExamplePostForm(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "Example Post with name: "+name+" and email: "+email)
}

// ExamplePostMultiPart ...
func ExamplePostMultiPart(c echo.Context) error {
	name := c.FormValue("name")
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	// Source
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")
}
