package approval

import (
	"context"
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
	// itemID := getItemIDfromEmbeds(i.Message.Embeds)

	client := mongorm.ConnectedClient()
	defer mongorm.DisconnectClient(client)

	// coll := client.Database(mongorm.DatabaseName).Collection(ToApprovalCollectionName)

	// var model *item.ItemModel

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

	var model *item.ItemModel
	err := model.Read(context.TODO(), coll, bson.M{"item.custom_id": itemID}, &model)
	if err != nil {
		log.Fatalf("Error on read item %v: %v", itemID, err)
	}
	log.Print(model)

	model.Delete(context.Background(), coll, bson.M{"item.custom_id": itemID})
}
