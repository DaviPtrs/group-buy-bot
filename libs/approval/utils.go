package approval

import (
	"github.com/bwmarrin/discordgo"
)

func getItemIDfromEmbeds(embeds []*discordgo.MessageEmbed) string {
	for _, e := range embeds {
		for _, field := range e.Fields {
			if field.Name == "ID" {
				return field.Value
			}
		}
	}
	return ""
}
