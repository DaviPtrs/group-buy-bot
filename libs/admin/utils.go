package admin

import (
	"log"
	"time"

	"github.com/DaviPtrs/group-buy-bot/libs/bot/session"
	"github.com/bwmarrin/discordgo"
)

func WrongChannelResponse(s *discordgo.Session, i *discordgo.Interaction, rightChannelID string) {
	message := "Parou a palhaçada ai, você não tem permissão pra usar esse comando!"
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
	time.Sleep(time.Second * 5)
	s.InteractionResponseDelete(i)
}

func GetUserName(id string) string {
	s := session.GetDiscordSession()
	guildID := session.GetGuildID()
	member, err := s.GuildMember(guildID, id)
	if err != nil {
		log.Fatalf("could not find user %s in guild %s", id, guildID)
	}
	return member.User.String()
}
