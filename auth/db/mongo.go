package db

import (
	"auth/config"
	"auth/internal/models"
	"context"

	"github.com/adityak368/swissknife/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBManager -- Struct To hold DB Connections
type dbManager struct {
	client     *mongo.Client
	connection *mongo.Database
}

var dbHandle dbManager

// MongoConnect -- Connect to DB with config
func MongoConnect() {

	dbHandle = dbManager{}

	//certPath, isCertificateAvailable := os.LookupEnv("MONGO_DB_CERT_PATH")

	clientOptions := options.Client().ApplyURI(config.DBUrl)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		logger.Info.Println("Failed To Connect To MongoDB")
		logger.Error.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		logger.Error.Println("Failed to ping MongoDB")
		logger.Error.Fatal(err)
	}

	models.PopulateIndex(client, config.DbName, "User")
	models.ListIndexes(client, config.DbName, "User")
	dbHandle.client = client
	dbHandle.connection = client.Database(config.DbName)
	logger.Info.Println("Connected to MongoDB!")
}

// GetMongoConnection -- Returns DB Handle
func GetMongoConnection() *mongo.Database {
	return dbHandle.connection
}

// MongoClose -- Close DB Connection
func MongoClose() {
	err := dbHandle.client.Disconnect(context.TODO())
	if err != nil {
		logger.Error.Println("Failed To Disconnect from MongoDB")
		logger.Error.Println(err)
	}
	logger.Info.Println("Connection to MongoDB closed.")
}
