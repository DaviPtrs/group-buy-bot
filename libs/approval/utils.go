package approval

import (
	"log"

	"github.com/DaviPtrs/group-buy-bot/libs/item"
	"github.com/DaviPtrs/group-buy-bot/libs/mongorm"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
)

func popFromToApproval(id string) *item.ItemModel {
	client := mongorm.ConnectedClient()
	defer mongorm.DisconnectClient(client)

	coll := client.Database(mongorm.DatabaseName).Collection(ToApprovalCollectionName)

	var model *item.ItemModel = new(item.ItemModel)
	err := model.Read(coll, bson.M{"item.custom_id": id}, model)
	if err != nil {
		log.Fatalf("Failed to find item %v: %v", id, err)
		return nil
	}

	err = model.Delete(coll, bson.M{"item.custom_id": id})
	if err != nil {
		log.Fatalf("Failed to remove item %v from to_approval list: %v", id, err)
	}
	return model
}

func pushToApproved(model *item.ItemModel) {
	client := mongorm.ConnectedClient()
	defer mongorm.DisconnectClient(client)

	coll := client.Database(mongorm.DatabaseName).Collection(ApprovedCollectionName)
	err := model.Create(coll, model)
	if err != nil {
		log.Fatalf("Failed to create approved item: %v", err)
	}
}

func getItemIDfromEmbeds(embeds []*discordgo.MessageEmbed) string {
	for _, e := range embeds {
		for _, field := range e.Fields {
			if field.Name == "ID" {
				return field.Value
			}
		}
	}
	return ""
}
