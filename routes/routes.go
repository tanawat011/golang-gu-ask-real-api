package routes

import (
	"github.com/labstack/echo"

	crud "golang-gu-ask-real-api/features/crud/controllers"
	"golang-gu-ask-real-api/features/users"
	"golang-gu-ask-real-api/services/mongodb"
)

// RouteCollect ...
func RouteCollect(e *echo.Echo, db *mongodb.Mongodb) {
	controller := crud.NewCRUD(db)

	g := e.Group("/oauth/v1")
	g.POST("/login", users.UserLogin)

	gCrud := e.Group("/crud")
	gCrud.GET("/get", controller.Get)
	gCrud.GET("/get-all", controller.GetAll)
	gCrud.POST("/create", controller.Create)
	// gCrud.PUT("/replace", controller.Replace)
	gCrud.PATCH("/update", controller.Update)
	gCrud.DELETE("/delete", controller.Delete)
}
