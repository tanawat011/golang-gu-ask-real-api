package main

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

// ExampleDataBody ...
type ExampleDataBody struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func main() {
	e := echo.New()
	e.GET("/:id", ExampleGet)
	e.POST("/json-body", ExamplePostJSONBody)
	e.POST("/xml-body", ExamplePostXMLBody)
	e.POST("/form-urlencoded", ExamplePostForm)
	e.POST("/multipart-form-data", ExamplePostMultiPart)
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
