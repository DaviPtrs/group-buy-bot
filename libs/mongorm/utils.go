package mongorm

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func ConnectedClient() *mongo.Client {
	client, err := ConnectWithURI(mongoURI)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Successfully connected to MongoDB")
	return client
}

func DisconnectClient(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Panicf("Error on disconnect from MongoDB: %v", err)
	}
}

func AddIndexes(coll *mongo.Collection, indexes []mongo.IndexModel) {
	_, err := coll.Indexes().CreateMany(context.TODO(), indexes)
	if err != nil {
		log.Fatalf("Failed to create indexes on collection %v: %v", coll.Name(), err)
	}
}
