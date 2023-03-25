package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Database

type MongoInstance struct {
	Client   *mongo.Client
	Database *mongo.Database
}

var MG MongoInstance

func Connect() {
	uri := os.Getenv("DATABASE_URI")
	database := os.Getenv("DATABASE_NAME")

	// Create a Client to a MongoDB server and use Ping to verify that the
	// server is running.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	db := client.Database(database)

	MG = MongoInstance{
		Client:   client,
		Database: db,
	}

	// Call Ping to verify that the deployment is up and the Client was
	// configured successfully. As mentioned in the Ping documentation, this
	// reduces application resiliency as the server may be temporarily
	// unavailable when Ping is called.
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected!")
}
