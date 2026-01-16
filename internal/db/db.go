package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectDB() *mongo.Database {
	const uri = "mongodb://localhost:27017/?directConnection=true"

	client, err := mongo.Connect(options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal("Cant connect to DB", err)
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}
	fmt.Println("Connected to MongoDB")

	coll := client.Database("mydb")

	fmt.Println("Connected to Database Successfully")

	return coll
}

type DB struct{
	Db *mongo.Database
}