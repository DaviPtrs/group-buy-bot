package user

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func CommandHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"add": addCommandHandler,
	}
}

func addCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	responseData := discordgo.InteractionResponseData{
		Content: "Add command invoked. Here's your response!",
		Flags:   discordgo.MessageFlagsEphemeral,
	}
	log.Print("handler called")
	response := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &responseData,
	}

	s.InteractionRespond(i.Interaction, &response)
}
