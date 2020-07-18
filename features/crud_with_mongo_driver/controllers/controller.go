package controllers

import (
	"golang-gu-ask-real-api/services/mongodb"

	"github.com/labstack/echo"
)

type ICRUD interface {
	Get(c echo.Context) error
	GetAll(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
}

type Tools struct {
	db *mongodb.Mongodb
}

func NewCRUD(db *mongodb.Mongodb) *Tools {
	return &Tools{
		db: db,
	}
}
