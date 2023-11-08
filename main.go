package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/DaviPtrs/group-buy-bot/libs/approval"
	"github.com/DaviPtrs/group-buy-bot/libs/user"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var discord *discordgo.Session
var guildID string
var removeCommands bool

func init() {
	godotenv.Load()
}

func init() {
	token, ok := os.LookupEnv("DISCORD_BOT_TOKEN")
	if !ok {
		log.Fatal("Bot Token not found")
	}
	guildID, ok = os.LookupEnv("DISCORD_BOT_GUILD_ID")
	if !ok {
		log.Fatal("Guild ID not found")
	}
	env, ok := os.LookupEnv("DISCORD_BOT_ENVIRONMENT")
	if ok && env == "development" {
		removeCommands = true
	}

	var err error
	discord, err = discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
}

func main() {
	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			if h, ok := user.CommandHandlers()[i.ApplicationCommandData().Name]; ok {
				h(s, i)
			}
		case discordgo.InteractionModalSubmit:
			for _, h := range user.ModalHandlers() {
				h(s, i)
			}
		case discordgo.InteractionMessageComponent:
			for _, h := range approval.ButtonHandlers() {
				h(s, i)
			}
			// data := i.MessageComponentData()
			// log.Print(data)
			// for _, e := range i.Message.Embeds {
			// 	for _, field := range e.Fields {
			// 		log.Print(field.Name)
			// 		log.Print(field.Value)
			// 	}
			// }
		}
	})

	err := discord.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	for _, v := range user.Commands {
		_, err := discord.ApplicationCommandCreate(discord.State.User.ID, guildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
	}

	defer discord.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	if removeCommands {
		addedCommands, err := discord.ApplicationCommands(discord.State.Application.ID, guildID)
		if err != nil {
			log.Printf("Failed to fetch commands %v", err)
		}

		log.Printf("Deleting commands from guild %v", guildID)
		for _, c := range addedCommands {
			err := discord.ApplicationCommandDelete(discord.State.User.ID, guildID, c.ID)
			if err != nil {
				log.Printf("Cannot delete '%v' command: %v", c.Name, err)
			}
		}

	}

	log.Println("Gracefully shutting down.")
}
