package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"piedpiper/utils/log"
)

var db *mongo.Database

// InitDatabase ...
func InitDatabase(mongoURI string, dbname string) error {
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return err
	}
	err = client.Connect(context.Background())
	if err != nil {
		return err
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return err
	}

	db = client.Database(dbname)
	log.Println("connected ok to the database")
	return nil
}
