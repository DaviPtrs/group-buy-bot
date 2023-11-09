package mongorm

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func ConnectedClient() *mongo.Client {
	client, err := ConnectWithURI(mongoURI)
	if err != nil {
		logrus.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	logrus.Infof("Successfully connected to MongoDB")
	return client
}

func DisconnectClient(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		logrus.Errorf("Error on disconnect from MongoDB: %v", err)
	}
}

func AddIndexes(coll *mongo.Collection, indexes []mongo.IndexModel) {
	_, err := coll.Indexes().CreateMany(context.TODO(), indexes)
	if err != nil {
		logrus.Errorf("Failed to create indexes on collection %v: %v", coll.Name(), err)
	}
}
