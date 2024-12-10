package db

import (
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FeatureFlag struct {
	Id        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Value     bool               `bson:"value" json:"value"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func (dbs *DatabaseService) InsertMockData() error {
	res, err := dbs.ffCol.InsertOne(dbs.ctx, bson.D{
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

func (dbs *DatabaseService) CreateFF(ffDTO *FFDto) (*mongo.InsertOneResult, error) {
	ff := FeatureFlag{
		Name:      ffDTO.Name,
		Value:     ffDTO.Value,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	res, err := dbs.ffCol.InsertOne(dbs.ctx, ff)

	if err != nil {
		return nil, err
	}

	return res, nil
}
