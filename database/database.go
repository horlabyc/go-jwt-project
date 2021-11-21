package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/horlabyc/golang-jwt-project/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBClient() *mongo.Client {
	uri := helpers.LoadEnv("MONGO_URI")
	if uri == "" {
		log.Fatal("No 'MONGODB_URI' environmental variable.")
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database")
	return client
}

var Client *mongo.Client = DBClient()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("go-auth").Collection(collectionName)
	return collection
}
