package database

import (
	"context"
	"fmt"
	"go-service/payx/configs"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {

	client, err := mongo.NewClient(options.Client().ApplyURI(configs.MongoDBEnvUrl()))
	if err != nil {
		log.Fatal(err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB Connected Successfully!")
	return client
}

var Client *mongo.Client = DBinstance()

func PayxCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var PayxCollection *mongo.Collection = client.Database("Payx").Collection(collectionName)
	return PayxCollection
}
