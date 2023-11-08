package approval

import (
	"fmt"
	"log"
	"os"

	"github.com/DaviPtrs/group-buy-bot/libs/item"
	"github.com/DaviPtrs/group-buy-bot/libs/mongorm"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var ApprovalChannelID string

func init() {
	godotenv.Load()

	var ok bool
	ApprovalChannelID, ok = os.LookupEnv("DISCORD_BOT_APPROVAL_CHANNEL_ID")
	if !ok {
		log.Fatal("Approval Channel ID not found")
	}
	seedDB()
}

func seedDB() {
	client := mongorm.ConnectedClient()

	var coll *mongo.Collection
	coll = client.Database(mongorm.DatabaseName).Collection("items_to_approval")
	mongorm.AddIndexes(coll, item.Indexes)

	coll = client.Database(mongorm.DatabaseName).Collection("items_approved")
	mongorm.AddIndexes(coll, item.Indexes)

	defer mongorm.DisconnectClient(client)
}

func SendItemToApproval(s *discordgo.Session, userID string, data *discordgo.ModalSubmitInteractionData) error {
	item, err := item.ParseFromModal(data)
	if err != nil {
		return err
	}
	log.Print(item)

	embed := discordgo.MessageEmbed{
		Fields: *item.ParseToEmbedFields(),
	}

	message := discordgo.MessageSend{
		Content: fmt.Sprintf("Item received from <@%s>\n\n", userID),
		Embeds:  []*discordgo.MessageEmbed{&embed},
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						CustomID: item.CustomID + "_approve_btn",
						Label:    "Approve",
						Style:    discordgo.PrimaryButton,
						Disabled: false,
					},
					discordgo.Button{
						CustomID: item.CustomID + "_reject_btn",
						Label:    "Reject",
						Style:    discordgo.DangerButton,
						Disabled: false,
					},
				},
			},
		},
	}
	_, err = s.ChannelMessageSendComplex(ApprovalChannelID, &message)

	if err != nil {
		log.Panicf("Error on sending item to approval: %v", err)
	}

	return nil
}
