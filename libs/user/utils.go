package user

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
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
		logrus.Errorf("Error on sending User \"Wrong Channel\" Message: %v", err)
		return
	}

	time.Sleep(time.Second * 10)
	err = s.InteractionResponseDelete(i)
	if err != nil {
		logrus.Errorf("Error on deleting User \"Wrong Channel\" response: %v", err)
	}
}
