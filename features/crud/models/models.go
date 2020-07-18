package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserCreate : Model for request create new user
type UserCreate struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// UserUpdate : Model for request update exiting user
type UserUpdate struct {
	Filter UserReceiver `json:"filter"`
	Update UserReceiver `json:"update"`
}

// UserFind : Model for request get one record
type UserFind struct {
	ID string `query:"id"`
}

// UserReceiver : Model for receive data
type UserReceiver struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	Name  string             `json:"name,omitempty" bson:"name"`
	Email string             `json:"email,omitempty" bson:"email"`
	Age   int                `json:"age,omitempty" bson:"age"`
}

// UserReceiverList : Model for receive list of data
type UserReceiverList []*UserReceiver
