package drivers

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var err error

type MongoDriver struct{}

func (md MongoDriver) GetMongoClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URL"))
	client, err := mongo.NewClient(clientOptions)
	log.Fatal(err)

	err = client.Connect(context.Background())
	log.Fatal(err)

	log.Println("Connected to Mongo")
	return client
}
