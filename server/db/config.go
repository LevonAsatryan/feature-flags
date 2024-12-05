package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	uri := os.Getenv("MONGODB_URI")
	log.Println("got mongodb url as: ", uri)

	// DB part
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable in .env file.")
	}

	client, err := mongo.Connect(ctx, options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return client
}
