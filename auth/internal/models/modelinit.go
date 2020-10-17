package models

import (
	"auth/config"
	"context"
	"fmt"
	"time"

	"github.com/adityak368/swissknife/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PopulateIndex(client *mongo.Client, database, collection string) {
	c := client.Database(database).Collection(collection)
	opts := options.CreateIndexes().SetMaxTime(time.Duration(config.DbTimeout) * time.Second)
	index := yieldUserIndexModel()
	c.Indexes().CreateOne(context.Background(), index, opts)
}

func yieldUserIndexModel() mongo.IndexModel {
	keys := bson.D{{"email", 1}}
	index := mongo.IndexModel{}
	index.Keys = keys
	index.Options = &options.IndexOptions{Unique: &[]bool{true}[0]}
	return index
}

func ListIndexes(client *mongo.Client, database, collection string) {
	c := client.Database(database).Collection(collection)
	duration := time.Duration(config.DbTimeout) * time.Second
	batchSize := int32(10)
	cur, err := c.Indexes().List(context.Background(), &options.ListIndexesOptions{&batchSize, &duration})
	if err != nil {
		logger.Error.Fatalf("Something went wrong listing %v", err)
	}
	for cur.Next(context.Background()) {
		index := bson.D{}
		cur.Decode(&index)
		logger.Info.Println(fmt.Sprintf("index found %v", index))
	}
}
