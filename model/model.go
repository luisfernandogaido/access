package model

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	uriMongo = "mongodb://gaido:1000sonhosreais@167.99.55.99:27017/?authSource=admin"
)

var (
	db *mongo.Database
)

func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI(uriMongo))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Connect(nil); err != nil {
		log.Fatal(err)
	}
	db = client.Database("php")
}
