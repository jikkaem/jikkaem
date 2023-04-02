package mongodb

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const database string = "mongogin-prod"

func GetColl(collName string) (*mongo.Collection, error) {
	client, err := GetMongoClient()
	return client.Database(database).Collection(collName), err
}

const uri = "mongodb://localhost:5000/jikkaem-userdb"

var MongoClient *mongo.Client
var MongoClientError error
var mongoOnce sync.Once

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {

		var serverAPI = options.ServerAPI(options.ServerAPIVersion1)
		var opts = options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

		client, err := mongo.Connect(context.TODO(), opts)

		MongoClient = client
		MongoClientError = err
	})
	return MongoClient, MongoClientError
}
