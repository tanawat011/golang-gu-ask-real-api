package models

import "github.com/Kamva/mgm"

type User struct {
	// DefaultModel add _id,created_at and updated_at fields to the Model
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Email            string `json:"email" bson:"email"`
	Age              int    `json:"age" bson:"age"`
}

func NewUser(name string, email string, age int) *User {
	return &User{
		Name:  name,
		Email: email,
		Age:   age,
	}
}
