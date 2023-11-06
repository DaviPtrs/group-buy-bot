package approval

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/DaviPtrs/group-buy-bot/libs/item"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var ApprovalChannelID string

func init() {
	godotenv.Load()

	var ok bool
	ApprovalChannelID, ok = os.LookupEnv("DISCORD_BOT_APPROVAL_CHANNEL_ID")
	if !ok {
		log.Fatal("Approval Channel ID not found")
	}
}

func SendItemToApproval(s *discordgo.Session, userID string, data *discordgo.ModalSubmitInteractionData) error {
	item, err := item.ParseFromModal(data)
	if err != nil {
		return err
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Item received from <@%s>\n\n", userID))
	sb.WriteString(fmt.Sprintf("**Link:** %s\n\n", item.URL))
	sb.WriteString(fmt.Sprintf("**Price:** $ %.2f\n\n", item.Price))
	sb.WriteString(fmt.Sprintf("**Weight:** %.2f lbs\n\n", item.Weight))
	sb.WriteString(fmt.Sprintf("**Estimated Tax:** %v %%\n\n", item.TaxRate))
	sb.WriteString(fmt.Sprintf("**Buyer's location:** %v\n\n", item.BuyerLocation))

	message := discordgo.MessageSend{
		Content: sb.String(),
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
