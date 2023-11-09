package mongorm

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoURI string
var DatabaseName string

func init() {
	godotenv.Load()
	var ok bool
	mongoURI, ok = os.LookupEnv("DISCORD_BOT_MONGODB_URI")
	if !ok {
		logrus.Fatal("Could not find MONGODB URI")
	}

	DatabaseName, ok = os.LookupEnv("DISCORD_BOT_MONGODB_DATABASE_NAME")
	if !ok {
		logrus.Fatal("Could not find Database Name")
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

	return client, nil
}
