package routes

import (
	"github.com/labstack/echo"

	crud "golang-gu-ask-real-api/features/crud_with_mongo_driver/controllers"
	crudWithMongoMgm "golang-gu-ask-real-api/features/crud_with_mongo_mgm/controllers"
	"golang-gu-ask-real-api/features/users"
	"golang-gu-ask-real-api/services/mongodb"
)

// RouteCollect ...
func RouteCollect(e *echo.Echo, db *mongodb.Mongodb) {
	cCrud := crud.NewCRUD(db)

	g := e.Group("/oauth/v1")
	g.POST("/login", users.UserLogin)

	gCrud := e.Group("/crud")
	gCrud.GET("/get", cCrud.Get)
	gCrud.GET("/get-all", cCrud.GetAll)
	gCrud.POST("/create", cCrud.Create)
	// gCrud.PUT("/replace", cCrud.Replace)
	gCrud.PATCH("/update", cCrud.Update)
	gCrud.DELETE("/delete", cCrud.Delete)

	gMgm := e.Group("/crudmgm")
	gMgm.GET("/getone", crudWithMongoMgm.GetOne)
	gMgm.GET("/get-by-filter", crudWithMongoMgm.GetByFilter)
}
