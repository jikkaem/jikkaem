package mongodb

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MongoClientError error
var mongoOnce sync.Once

func GetMongoClient(uri string) (*mongo.Client, error) {
	mongoOnce.Do(func() {

		var serverAPI = options.ServerAPI(options.ServerAPIVersion1)
		var opts = options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

		client, err := mongo.Connect(context.TODO(), opts)

		MongoClient = client
		MongoClientError = err
	})
	return MongoClient, MongoClientError
}
