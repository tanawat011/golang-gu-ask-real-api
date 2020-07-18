package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// IMongodb : Interface mongodb service in app
type IMongodb interface {
	MongoConnect() *mongo.Client
	InsertOne(collection string, data interface{}) error
	UpdateOne(collection string, filter bson.D, data bson.D) error
	FindOne(collection string, filter primitive.D, receiver interface{}) error
	Find(collection string, filter primitive.D, receiver interface{}) error
	Delete(collection string, filter primitive.D) error
}

// Mongodb : Connector model
type Mongodb struct {
	connector *mongo.Client
	database  *mongo.Database
	ctx       context.Context
}

var uri = "mongodb://localhost:27017"
var db = "example_crud"

// NewConnectionMongo : Create new connection mongodb
func NewConnectionMongo(ctx context.Context) *Mongodb {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	if err = client.Connect(ctx); err != nil {
		panic(err)
	}

	// Check the connection
	if err = client.Ping(ctx, nil); err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")

	// Connect to the specified database
	db := client.Database(db)

	return &Mongodb{
		connector: client,
		database:  db,
		ctx:       ctx,
	}
}

// MongoConnect : Call mongodb client
func (m *Mongodb) MongoConnect() *mongo.Client {
	return m.connector
}

// InsertOne ...
func (m *Mongodb) InsertOne(collection string, data interface{}) error {
	client := m.connector
	ctx := m.ctx

	db := client.Database(db)
	c := db.Collection(collection)

	result, err := c.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	fmt.Println("Inserted a single document: ", result.InsertedID)

	return nil
}

// UpdateOne ...
func (m *Mongodb) UpdateOne(collection string, filter bson.D, data bson.D) error {
	client := m.connector
	ctx := m.ctx

	db := client.Database(db)
	c := db.Collection(collection)

	result, err := c.UpdateOne(ctx, filter, data)
	if err != nil {
		return err
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", result.MatchedCount, result.ModifiedCount)

	return nil
}

// FindOne ...
func (m *Mongodb) FindOne(collection string, filter primitive.D, receiver interface{}) error {
	client := m.connector
	ctx := m.ctx

	db := client.Database(db)
	c := db.Collection(collection)

	if err := c.FindOne(ctx, filter).Decode(receiver); err != nil {
		return err
	}
	fmt.Printf("Found a single document: %+v\n", receiver)

	return nil
}

// Find ...
func (m *Mongodb) Find(collection string, filter primitive.D, receiver *[]map[string]interface{}) error {
	ctx := m.ctx
	db := m.database
	col := db.Collection(collection)

	findOptions := options.Find()
	cur, err := col.Find(ctx, filter, findOptions)
	if err != nil {
		return err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var elem map[string]interface{}
		if err := cur.Decode(&elem); err != nil {
			return err
		}
		*receiver = append(*receiver, elem)
	}
	if err := cur.Err(); err != nil {
		return err
	}
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", receiver)

	return nil
}

// Delete ...
func (m *Mongodb) Delete(collection string, filter primitive.D) error {
	ctx := m.ctx
	db := m.database
	col := db.Collection(collection)

	deleteResult, err := col.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	return nil
}
