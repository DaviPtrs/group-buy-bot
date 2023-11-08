package user

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func WrongChannelResponse(s *discordgo.Session, i *discordgo.Interaction, rightChannelID string) {
	message := fmt.Sprintf("Canal errado, amigo. Esse bot só responde comandos no <#%v>", rightChannelID)
	err := s.InteractionRespond(i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 10)
	s.InteractionResponseDelete(i)
}
