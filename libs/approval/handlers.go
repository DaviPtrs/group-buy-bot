package approval

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
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
	model := popFromToApproval(itemID)
	pushToApproved(model)

	submit_message := "Item approved!"
	embed := discordgo.MessageEmbed{
		Fields: *model.Item.ParseToEmbedFields(),
	}
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Embeds:  []*discordgo.MessageEmbed{&embed},
			Content: submit_message,
		},
	})
	if err != nil {
		log.Panicf("Unable to respond to modal %v: %v", data.CustomID, err)
	}

	SendItemFeedback(s, &model.Item, nil)
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

func rejectModalHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ModalSubmitData()
	if !strings.HasPrefix(data.CustomID, "rejected_") {
		return
	}

	itemID := strings.TrimPrefix(data.CustomID, "rejected_")
	model := popFromToApproval(itemID)

	reason := data.Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value
	submit_message := fmt.Sprintf("Item rejected! Reason: %v", reason)
	embed := discordgo.MessageEmbed{
		Fields: *model.Item.ParseToEmbedFields(),
	}
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Embeds:  []*discordgo.MessageEmbed{&embed},
			Content: submit_message,
		},
	})
	if err != nil {
		log.Panicf("Unable to respond to modal %v: %v", data.CustomID, err)
	}

	SendItemFeedback(s, &model.Item, &reason)
}

func rejectItemHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.MessageComponentData()
	if !(strings.HasPrefix(data.CustomID, "group_buy_item_") && strings.HasSuffix(data.CustomID, "_reject_btn")) {
		return
	}

	itemID := getItemIDfromEmbeds(i.Message.Embeds)
	if itemID == "" {
		log.Fatal("There's no item in this message")
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
