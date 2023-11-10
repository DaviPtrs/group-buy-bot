package item

import (
	"github.com/DaviPtrs/group-buy-bot/libs/mongorm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ToApprovalCollectionName = "items_to_approval"
var ApprovedCollectionName = "items_approved"

var trueVar = true // why???? WHY DO I HAVE TO DO THIS?

var Indexes = []mongo.IndexModel{
	{
		Keys:    bson.D{{Key: "item.custom_id", Value: 1}},
		Options: &options.IndexOptions{Unique: &trueVar},
	},
	{
		Keys: bson.D{{Key: "item.user_id", Value: 1}},
	},
}

type ItemModel struct {
	mongorm.Model
	Item
}

func SeedDB() {
	client := mongorm.ConnectedClient()

	var coll *mongo.Collection
	coll = client.Database(mongorm.DatabaseName).Collection(ToApprovalCollectionName)
	mongorm.AddIndexes(coll, Indexes)

	coll = client.Database(mongorm.DatabaseName).Collection(ApprovedCollectionName)
	mongorm.AddIndexes(coll, Indexes)

	defer mongorm.DisconnectClient(client)
}
