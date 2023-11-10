package approval

import (
	"fmt"
	"os"

	"github.com/DaviPtrs/group-buy-bot/libs/item"
	"github.com/DaviPtrs/group-buy-bot/libs/mongorm"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var ApprovalChannelID string

func init() {
	godotenv.Load()

	var ok bool
	ApprovalChannelID, ok = os.LookupEnv("DISCORD_BOT_APPROVAL_CHANNEL_ID")
	if !ok {
		logrus.Fatal("Approval Channel ID not found")
	}
	item.SeedDB()
}

func SendItemToApproval(s *discordgo.Session, userID string, data *discordgo.ModalSubmitInteractionData) error {
	i, err := item.ParseFromModal(data)
	if err != nil {
		return err
	}
	// logrus.Print(i)

	embed := discordgo.MessageEmbed{
		Fields: *i.ParseToEmbedFields(),
	}

	message := discordgo.MessageSend{
		Content: fmt.Sprintf("Item received from <@%s>\n\n", userID),
		Embeds:  []*discordgo.MessageEmbed{&embed},
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						CustomID: i.CustomID + "_approve_btn",
						Label:    "Approve",
						Style:    discordgo.PrimaryButton,
						Disabled: false,
					},
					discordgo.Button{
						CustomID: i.CustomID + "_reject_btn",
						Label:    "Reject",
						Style:    discordgo.DangerButton,
						Disabled: false,
					},
				},
			},
		},
	}
	model := i.GetModel()

	client := mongorm.ConnectedClient()
	defer mongorm.DisconnectClient(client)

	coll := client.Database(mongorm.DatabaseName).Collection(item.ToApprovalCollectionName)

	err = model.Create(coll, model)
	if err != nil {
		return fmt.Errorf("failed to create to_approval item: %v", err)
	}

	_, err = s.ChannelMessageSendComplex(ApprovalChannelID, &message)

	if err != nil {
		return fmt.Errorf("error on sending item to approval channel: %v", err)
	}

	return nil
}
