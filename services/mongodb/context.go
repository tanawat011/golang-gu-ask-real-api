package mongodb

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db = "example_crud"

// InsertOne ...
func InsertOne(collection string, data interface{}) error {
	client, ctx, err := mongoConnect()
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)

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
func UpdateOne(collection string, filter bson.D, data bson.D) error {
	client, ctx, err := mongoConnect()
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)

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
func FindOne(collection string, filter primitive.D, receiver interface{}) error {
	client, ctx, err := mongoConnect()
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)

	db := client.Database(db)
	c := db.Collection(collection)

	err = c.FindOne(ctx, filter).Decode(receiver)
	if err != nil {
		return err
	}
	fmt.Printf("Found a single document: %+v\n", receiver)

	return nil
}

// Find ...
func Find(collection string, filter primitive.D, receiver interface{}) error {
	client, ctx, err := mongoConnect()
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)

	db := client.Database(db)
	c := db.Collection(collection)

	findOptions := options.Find()
	findOptions.SetLimit(2)
	cur, err := c.Find(ctx, filter, findOptions)
	if err != nil {
		return err
	}

	results := make([]map[string]interface{}, 0)
	for cur.Next(ctx) {
		var elem map[string]interface{}
		err := cur.Decode(&elem)
		if err != nil {
			return err
		}

		results = append(results, elem)
	}
	if err := cur.Err(); err != nil {
		return err
	}
	cur.Close(ctx)
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

	return nil
}
