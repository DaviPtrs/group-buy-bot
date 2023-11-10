package admin

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var AdminChannelID string

func init() {
	godotenv.Load()

	var ok bool
	AdminChannelID, ok = os.LookupEnv("DISCORD_BOT_ADMIN_CHANNEL_ID")
	if !ok {
		logrus.Fatal("Admin Channel ID not found")
	}
}

func CommandHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"dump":   dumpCommandHandler,
		"finish": finishCommandHandler,
	}
}

func dumpCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ChannelID != AdminChannelID {
		WrongChannelResponse(s, i.Interaction)
		return
	}

	response := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: nil,
	}
	err := s.InteractionRespond(i.Interaction, &response)
	if err != nil {
		logrus.Errorf("Error on responding \"%v\": %v", i.ApplicationCommandData().Name, err)
		return
	}

	modelList, err := FetchAllItems()
	if err != nil {
		logrus.Errorf("Error on fetching approved items from DB: %v", err)
		return
	}

	bufBytes, err := GenerateCSV(modelList)
	if err != nil {
		logrus.Errorf("Error on generating dump CSV: %v", err)
		return
	}

	fileName := fmt.Sprintf("group-buy-dump-%s.csv", time.Now().Format("2017-09-07-1504"))
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

	_, err = s.FollowupMessageCreate(i.Interaction, true, &options)
	if err != nil {
		logrus.Errorf("Error on sending dump message: %v", err)
	}
}

func finishCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ChannelID != AdminChannelID {
		WrongChannelResponse(s, i.Interaction)
		return
	}

	dumpCommandHandler(s, i)

	if err := dropApprovedItems(); err != nil {
		logrus.Errorf("Error on finishing group buy batch: %v", err)
	}

	options := discordgo.WebhookParams{
		Content: "Approved items marked as shipped and removed!",
	}
	_, err := s.FollowupMessageCreate(i.Interaction, true, &options)
	if err != nil {
		logrus.Errorf("Error on sending finish confirmation message: %v", err)
	}
}
