package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
var CTX context.Context

func configDB(ctx context.Context) (*mongo.Database, error) {
	// Creating new mongo client and passing connection string (getting string from env variable)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@mongodb:27017/?maxPoolSize=20"))
	if err != nil {
		log.Fatal("connection error: ", err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDb")

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("ping err: ", err)
	}

	// Using database set in env variable
	reviewsDb := client.Database("omnisend")

	// Returning pointer to a database
	return reviewsDb, nil
}

func init() {
	var err error
	// The mongo.Database initialization process requires a context.Context object
	CTX = context.Background()
	CTX, cancel := context.WithCancel(CTX)
	defer cancel()

	// Calling DB connection function giving context object
	DB, err = configDB(CTX)
	if err != nil {
		log.Fatal(err)
	}
}
