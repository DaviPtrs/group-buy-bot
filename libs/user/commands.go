package user

import "github.com/bwmarrin/discordgo"

var (
	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "add",
			Description: "Register/Add new product",
			DescriptionLocalizations: &map[discordgo.Locale]string{
				discordgo.PortugueseBR: "Registrar/Adicionar novo produto",
			},
		},
	}
)
