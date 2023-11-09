package admin

import (
	"time"

	"github.com/DaviPtrs/group-buy-bot/libs/bot/session"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func WrongChannelResponse(s *discordgo.Session, i *discordgo.Interaction) {
	message := "Parou a palhaçada ai, você não tem permissão pra usar esse comando!"
	err := s.InteractionRespond(i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		logrus.Errorf("Error on sending Admin \"Wrong Channel\" Message: %v", err)
		return
	}

	time.Sleep(time.Second * 5)
	err = s.InteractionResponseDelete(i)
	if err != nil {
		logrus.Errorf("Error on deleting Admin \"Wrong Channel\" response: %v", err)
	}
}

func GetUserName(id string) string {
	s := session.GetDiscordSession()
	guildID := session.GetGuildID()
	member, err := s.GuildMember(guildID, id)
	if err != nil {
		logrus.Errorf("could not find user %s in guild %s", id, guildID)
		return ""
	}
	return member.User.String()
}
