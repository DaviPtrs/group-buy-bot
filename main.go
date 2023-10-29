package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var discord *discordgo.Session

func init() {
	godotenv.Load()
}

func init() {
	var err error
	token, ok := os.LookupEnv("DISCORD_BOT_TOKEN")
	if !ok {
		log.Fatal("Bot Token not found")
	}
	discord, err = discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
}

func test(s *discordgo.Session, c *discordgo.Connect) {
	for _, guild := range s.State.Guilds {
		fmt.Println(guild.ID)
	}
}

func main() {
	discord.AddHandler(test)

	err := discord.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
	defer discord.Close()
}
