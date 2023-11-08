package approval

import (
	"log"
	"strings"

	"github.com/DaviPtrs/group-buy-bot/libs/item"
	"github.com/DaviPtrs/group-buy-bot/libs/mongorm"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
)

func ButtonHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"Approve": approveItemHandler,
		"Reject":  rejectItemHandler,
	}
}

func approveItemHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.MessageComponentData()
	if !(strings.HasPrefix(data.CustomID, "group_buy_item_") && strings.HasSuffix(data.CustomID, "_approve_btn")) {
		return
	}
	itemID := getItemIDfromEmbeds(i.Message.Embeds)
	if itemID == "" {
		log.Fatal("There's no item in this message")
	}

	client := mongorm.ConnectedClient()
	defer mongorm.DisconnectClient(client)

	coll := client.Database(mongorm.DatabaseName).Collection(ToApprovalCollectionName)

	var model *item.ItemModel = new(item.ItemModel)
	err := model.Read(coll, bson.M{"item.custom_id": itemID}, model)
	if err != nil {
		log.Fatalf("Failed to find item %v: %v", itemID, err)
		return
	}

	err = model.Delete(coll, bson.M{"item.custom_id": itemID})
	if err != nil {
		log.Fatalf("Failed to remove item %v from to_approval list: %v", itemID, err)
	}

	coll = client.Database(mongorm.DatabaseName).Collection(ApprovedCollectionName)
	err = model.Create(coll, model)
	if err != nil {
		log.Fatalf("Failed to create approved item: %v", err)
	}

}

func rejectItemHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.MessageComponentData()
	if !(strings.HasPrefix(data.CustomID, "group_buy_item_") && strings.HasSuffix(data.CustomID, "_reject_btn")) {
		return
	}
	itemID := getItemIDfromEmbeds(i.Message.Embeds)

	client := mongorm.ConnectedClient()
	defer mongorm.DisconnectClient(client)

	coll := client.Database(mongorm.DatabaseName).Collection(ToApprovalCollectionName)

	var model *item.ItemModel = new(item.ItemModel)
	err := model.Delete(coll, bson.M{"item.custom_id": itemID})

	if err != nil {
		log.Fatalf("Failed to remove item %v from to_approval list: %v", itemID, err)
	}

}
