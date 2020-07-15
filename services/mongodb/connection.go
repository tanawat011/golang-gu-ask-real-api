package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uri = "mongodb://localhost:27017"

// mongoConnect ...
func mongoConnect() (*mongo.Client, context.Context, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, nil, err
	}
	ctx := context.TODO()
	if err = client.Connect(ctx); err != nil {
		return nil, nil, err
	}

	// Check the connection
	if err = client.Ping(ctx, nil); err != nil {
		return nil, nil, err
	}
	fmt.Println("Connected to MongoDB!")

	return client, ctx, nil
}
