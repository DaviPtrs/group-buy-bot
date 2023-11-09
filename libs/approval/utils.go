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

func SendItemFeedback(s *discordgo.Session, i *item.Item, reasoning *string) {
	var sb strings.Builder
	if reasoning == nil {
		sb.WriteString("**Seu item foi APROVADO para o Group buy!**\n\n")
	} else {
		sb.WriteString("**Seu item foi REPROVADO para o Group buy!**\n\n")
		sb.WriteString(fmt.Sprintf("**Motivo:** %s\n", *reasoning))
		sb.WriteString("\nEm caso de discordancia sobre o motivo da recusa, entre o contato com os administradores do servidor.\n")
	}
	sb.WriteString("Caso tenha duvidas sobre o processo, leia as regras presentes no chat do Group Buy bot.\n")
	sb.WriteString("**Lembrete:** Items não elegíveis para o Group Buy podem (mediante análise) ser importados individualmente mediante a pedido especial.\n")
	sb.WriteString(fmt.Sprintf("\n**Item URL:** %v\n", i.URL))

	st, err := s.UserChannelCreate(i.UserID)
	if err != nil {
		log.Fatalf("Failed to create DM with user %v: %v", i.UserID, err)
	}
	s.ChannelMessageSend(st.ID, sb.String())
}

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
