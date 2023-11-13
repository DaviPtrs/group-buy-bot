package approval

import (
	"fmt"
	"strings"

	"github.com/DaviPtrs/group-buy-bot/libs/check"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func ButtonHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"Approve": approveItemHandler,
		"Reject":  rejectItemHandler,
	}
}

func ModalHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"Reject":  rejectModalHandler,
		"Approve": approveModalHandler,
	}
}

func approveModalHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ModalSubmitData()
	if !strings.HasPrefix(data.CustomID, "approved_") {
		return
	}

	itemID := strings.TrimPrefix(data.CustomID, "approved_")
	model, err := popFromToApproval(itemID)
	if err != nil {
		logrus.Error(err)
		return
	}
	if err := pushToApproved(model); err != nil {
		logrus.Errorf("Failed to create approved item: %v", err)
		return
	}

	submit_message := fmt.Sprintf("Item approved by <@%s>!", i.Interaction.Member.User.ID)
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
		logrus.Errorf("Unable to respond to modal %v: %v", data.CustomID, err)
		return
	}

	if err := SendItemFeedback(s, &model.Item, nil); err != nil {
		logrus.Error(err)
	}
	check.CheckReadyToBuy()

}

func approveItemHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.MessageComponentData()
	if !(strings.HasPrefix(data.CustomID, "group_buy_item_") && strings.HasSuffix(data.CustomID, "_approve_btn")) {
		return
	}

	itemID := getItemIDfromEmbeds(i.Message.Embeds)
	if itemID == "" {
		logrus.Errorf("Unable to fetch item %v from message", itemID)
		return
	}

	responseData := discordgo.InteractionResponseData{
		CustomID: "approved_" + itemID,
		Title:    "Do you really want to approve this crap?",
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.TextInput{
						CustomID:    "useless",
						Style:       discordgo.TextInputParagraph,
						Placeholder: "Useless field coz I can't create modals without components. So yeah, blame Discord",
						MaxLength:   1,
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
		logrus.Errorf("Failed to send response modal %v: %v", responseData.CustomID, err)
	}
}

func rejectModalHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ModalSubmitData()
	if !strings.HasPrefix(data.CustomID, "rejected_") {
		return
	}

	itemID := strings.TrimPrefix(data.CustomID, "rejected_")
	model, err := popFromToApproval(itemID)
	if err != nil {
		logrus.Error(err)
		return
	}

	reason := data.Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value
	submit_message := fmt.Sprintf("Item rejected by <@%s>! Reason: %v", i.Interaction.Member.User.ID, reason)
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
		logrus.Errorf("Unable to respond to modal %v: %v", data.CustomID, err)
		return
	}

	if err := SendItemFeedback(s, &model.Item, &reason); err != nil {
		logrus.Error(err)
	}
}

func rejectItemHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.MessageComponentData()
	if !(strings.HasPrefix(data.CustomID, "group_buy_item_") && strings.HasSuffix(data.CustomID, "_reject_btn")) {
		return
	}

	itemID := getItemIDfromEmbeds(i.Message.Embeds)
	if itemID == "" {
		logrus.Errorf("Unable to fetch item %v from message", itemID)
		return
	}

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
						MaxLength:   4000,
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
		logrus.Errorf("Failed to send response modal %v: %v", responseData.CustomID, err)
	}
}
