package user

import "github.com/bwmarrin/discordgo"

var (
	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "new",
			Description: "Register new product",
			DescriptionLocalizations: &map[discordgo.Locale]string{
				discordgo.PortugueseBR: "Registrar novo produto",
			},
		},
	}
)
