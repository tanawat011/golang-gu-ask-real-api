package mongodb

import (
	"log"

	"github.com/Kamva/mgm"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	uri    = "mongodb://localhost:27017"
	dbName = "example_crud"
)

// NewConnectionMongo : Create new connection mongodb
func NewConnectionMongo() {
	if err := mgm.SetDefaultConfig(nil, dbName, options.Client().ApplyURI(uri)); err != nil {
		panic(err)
	}
	log.Println("Connected to MongoDB!")
}
