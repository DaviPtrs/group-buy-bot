package admin

import "github.com/bwmarrin/discordgo"

var (
	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "dump",
			Description: "Dump Approved GB items to CSV",
		},
		// {
		// 	Name:        "del",
		// 	Description: "Dump Approved GB items to CSV",
		// },
	}
)
