package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const FF_TABLE_NAME = "ffs"
const DATABASE_NAME = "feature_flags"

type DatabaseService struct {
	client *mongo.Client
	ffCol  *mongo.Collection
	ctx    context.Context
}

func ConnectDB() *DatabaseService {
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

	ffCol := client.Database(DATABASE_NAME).Collection(FF_TABLE_NAME)

	dbService := DatabaseService{client: client, ffCol: ffCol, ctx: ctx}

	return &dbService
}
