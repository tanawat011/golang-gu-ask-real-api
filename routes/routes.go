package routes

import (
	"github.com/labstack/echo"

	"golang-gu-ask-real-api/features/crud"
	"golang-gu-ask-real-api/features/users"
)

// RouteCollect ...
func RouteCollect(e *echo.Echo) {
	g := e.Group("/oauth/v1")
	g.POST("/login", users.UserLogin)

	g2 := e.Group("/crud")
	g2.GET("/get", crud.Get)
	g2.GET("/get-all", crud.GetAll)
	g2.POST("/create", crud.Create)
	g2.PUT("/replace", crud.Replace)
	g2.PATCH("/update", crud.Update)
	g2.DELETE("/delete", crud.Delete)
}
