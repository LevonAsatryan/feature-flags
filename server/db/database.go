package db

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const FF_TABLE_NAME = "ffs"
const DATABASE_NAME = "feature_flags"

func InsertMockData(client *mongo.Client) error {
	coll := client.Database(DATABASE_NAME).Collection(FF_TABLE_NAME)
	res, err := coll.InsertOne(context.TODO(), bson.D{
		{"name", "feature 1"},
		{"value", true},
	})

	/*
	 * Replace this in the future with a logger
	 */
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", FF_TABLE_NAME)
		return err
	}

	if err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(res, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)

	return err
}
