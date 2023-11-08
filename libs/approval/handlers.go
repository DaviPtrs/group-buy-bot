package approval

import (
	"fmt"
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

func ModalHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"Reject": rejectModalHandler,
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

func rejectModalHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ModalSubmitData()
	if !strings.HasPrefix(data.CustomID, "rejected_") {
		return
	}
	itemID := strings.TrimPrefix(data.CustomID, "rejected_")

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

	reason := data.Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value
	submit_message := fmt.Sprintf("Item rejected! Reason: %v", reason)
	embed := discordgo.MessageEmbed{
		Fields: *model.Item.ParseToEmbedFields(),
	}
	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Embeds:  []*discordgo.MessageEmbed{&embed},
			Content: submit_message,
		},
	})
	if err != nil {
		log.Panicf("Unable to respond to modal %v: %v", data.CustomID, err)
	}
}

func rejectItemHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.MessageComponentData()
	if !(strings.HasPrefix(data.CustomID, "group_buy_item_") && strings.HasSuffix(data.CustomID, "_reject_btn")) {
		return
	}
	itemID := getItemIDfromEmbeds(i.Message.Embeds)

	responseData := discordgo.InteractionResponseData{
		CustomID: "rejected_" + itemID,
		Title:    "Reject item - Reason",
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.TextInput{
						CustomID:    "reject-reason",
						Label:       "Why is this item being rejected?",
						Style:       discordgo.TextInputParagraph,
						Placeholder: "Too heavy product",
						Required:    true,
					},
				},
			},
		},
	}
	response := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &responseData,
	}

	err := s.InteractionRespond(i.Interaction, &response)
	if err != nil {
		log.Fatal(err)
	}

}
