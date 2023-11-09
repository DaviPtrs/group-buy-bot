package admin

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var AdminChannelID string

func init() {
	godotenv.Load()

	var ok bool
	AdminChannelID, ok = os.LookupEnv("DISCORD_BOT_ADMIN_CHANNEL_ID")
	if !ok {
		log.Fatal("Admin Channel ID not found")
	}
}

func CommandHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"dump": dumpCommandHandler,
	}
}

func dumpCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ChannelID != AdminChannelID {
		WrongChannelResponse(s, i.Interaction, AdminChannelID)
		return
	}

	response := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: nil,
	}
	err := s.InteractionRespond(i.Interaction, &response)
	if err != nil {
		log.Panicf("Error on responding \"%v\": %v", i.ApplicationCommandData().Name, err)
	}

	modelList := FetchAllItems()
	bufBytes := GenerateCSV(modelList)
	fileName := fmt.Sprintf("group-buy-dump-%s.csv", time.Now().Format("2017-09-07"))
	fileInfo := discordgo.File{
		Name:        fileName,
		ContentType: "multipart/form-data",
		Reader:      bytes.NewReader(bufBytes),
	}

	options := discordgo.WebhookParams{
		Files: []*discordgo.File{
			&fileInfo,
		},
	}
	s.FollowupMessageCreate(i.Interaction, true, &options)
}
