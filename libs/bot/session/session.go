package session

import (
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var discord *discordgo.Session
var guildID string

func init() {
	godotenv.Load()

	token, ok := os.LookupEnv("DISCORD_BOT_TOKEN")
	if !ok {
		logrus.Fatal("Bot Token not found")
	}
	guildID, ok = os.LookupEnv("DISCORD_BOT_GUILD_ID")
	if !ok {
		logrus.Fatal("Guild ID not found")
	}

	var err error
	discord, err = discordgo.New("Bot " + token)
	if err != nil {
		logrus.Fatalf("Failed to create session: %v", err)
	}
}

func GetDiscordSession() *discordgo.Session {
	return discord
}

func GetGuildID() string {
	return guildID
}
