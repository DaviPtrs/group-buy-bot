package bot

import (
	"os"
	"os/signal"

	"github.com/DaviPtrs/group-buy-bot/libs/admin"
	"github.com/DaviPtrs/group-buy-bot/libs/approval"
	"github.com/DaviPtrs/group-buy-bot/libs/bot/session"
	"github.com/DaviPtrs/group-buy-bot/libs/user"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var discord *discordgo.Session
var guildID string
var removeCommands bool

func init() {
	godotenv.Load()

	env, ok := os.LookupEnv("DISCORD_BOT_ENVIRONMENT")
	if ok && env == "development" {
		removeCommands = true
	}

	discord = session.GetDiscordSession()
	guildID = session.GetGuildID()
}

func interactionCreateHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		if h, ok := user.CommandHandlers()[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
		if h, ok := admin.CommandHandlers()[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	case discordgo.InteractionModalSubmit:
		for _, h := range user.ModalHandlers() {
			h(s, i)
		}
		for _, h := range approval.ModalHandlers() {
			h(s, i)
		}
	case discordgo.InteractionMessageComponent:
		for _, h := range approval.ButtonHandlers() {
			h(s, i)
		}
	}
}

func registerCommands(s *discordgo.Session) {
	for _, v := range user.Commands {
		_, err := discord.ApplicationCommandCreate(discord.State.User.ID, guildID, v)
		if err != nil {
			logrus.Errorf("Cannot create '%v' command: %v", v.Name, err)
		}
	}

	for _, v := range admin.Commands {
		_, err := discord.ApplicationCommandCreate(discord.State.User.ID, guildID, v)
		if err != nil {
			logrus.Errorf("Cannot create '%v' command: %v", v.Name, err)
		}
	}
}

func Run() {
	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		logrus.Infof("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	discord.AddHandler(interactionCreateHandler)

	err := discord.Open()
	if err != nil {
		logrus.Panicf("Cannot open the session: %v", err)
	}

	registerCommands(discord)

	defer discord.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	if removeCommands {
		addedCommands, err := discord.ApplicationCommands(discord.State.Application.ID, guildID)
		if err != nil {
			logrus.Errorf("Failed to fetch commands %v", err)
			return
		}

		logrus.Infof("Deleting commands from guild %v", guildID)
		for _, c := range addedCommands {
			err := discord.ApplicationCommandDelete(discord.State.User.ID, guildID, c.ID)
			if err != nil {
				logrus.Errorf("Cannot delete '%v' command: %v", c.Name, err)
			}
		}

	}

	logrus.Infof("Gracefully shutting down.")
}
