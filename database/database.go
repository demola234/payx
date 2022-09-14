package database

import (
	"context"
	"fmt"
	// "go-service/payx/configs"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	var uri = "mongodb+srv://payx:vU5l1aNBxBKc6Vhb@cluster0.bu7aw34.mongodb.net/?retryWrites=true&w=majority"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
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
