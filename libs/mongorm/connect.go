package mongorm

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoURI string

func init() {
	godotenv.Load()
	var ok bool
	mongoURI, ok = os.LookupEnv("DISCORD_BOT_MONGODB_URI")
	if !ok {
		log.Fatal("Could not find MONGODB URI")
	}
}

func ConnectWithURI(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	log.Print("Successfully connected to MongoDB")
	return client, nil
}

func ConnectedClient() *mongo.Client {
	client, err := ConnectWithURI(mongoURI)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
