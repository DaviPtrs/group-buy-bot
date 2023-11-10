package admin

import "github.com/bwmarrin/discordgo"

var (
	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "dump",
			Description: "Dump Approved GB items to CSV",
		},
		{
			Name:        "finish",
			Description: "Mark all approved items as shipped. This action is destructive.",
		},
		// {
		// 	Name:        "del",
		// 	Description: "Dump Approved GB items to CSV",
		// },
	}
)
