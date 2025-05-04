package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var BookCollection *mongo.Collection

func Connect() *mongo.Collection {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	log.Println("Connected to MongoDB")

	return client.Database("library").Collection("books")

}
