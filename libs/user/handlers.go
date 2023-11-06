package user

import (
	"github.com/bwmarrin/discordgo"
	uuid "github.com/satori/go.uuid"
)

func CommandHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"add": addCommandHandler,
	}
}

// func addCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
// 	responseData := discordgo.InteractionResponseData{
// 		Content: "Add command invoked. Here's your response!",
// 		Flags:   discordgo.MessageFlagsEphemeral,
// 	}
// 	log.Print("handler called")
// 	response := discordgo.InteractionResponse{
// 		Type: discordgo.InteractionResponseChannelMessageWithSource,
// 		Data: &responseData,
// 	}

// 	s.InteractionRespond(i.Interaction, &response)
// }

func addCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	uuid := uuid.NewV4()
	responseData := discordgo.InteractionResponseData{
		CustomID: "group_buy_item_" + i.Interaction.Member.User.ID + "_" + uuid.String(),
		// Flags:    discordgo.MessageFlagsEphemeral,
		Title: "Group buy - Request new product",
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.TextInput{
						CustomID:    "product-url",
						Label:       "Link do produto",
						Style:       discordgo.TextInputShort,
						Placeholder: "https://www.amazon.com/Images-You-Should-Not-Masturbate/dp/0399536493",
						Required:    true,
					},
				},
			},
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.TextInput{
						CustomID:    "price",
						Label:       "Preço (incluindo o frete, se tiver)",
						Placeholder: "Preço total em dólar ($)",
						Style:       discordgo.TextInputShort,
						Required:    true,
					},
				},
			},
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.TextInput{
						CustomID:    "weight",
						Label:       "Peso total (não esqueça da unidade de medida)",
						Placeholder: "0.5kg ou 1lb",
						Style:       discordgo.TextInputShort,
						Required:    false,
					},
				},
			},
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.TextInput{
						CustomID:    "tax-rate",
						Label:       "Quanto você está disposto a pagar de imposto?",
						Placeholder: "Taxa em porcentagem (0% não vale). Ex: 20%",
						Style:       discordgo.TextInputShort,
						Required:    true,
					},
				},
			},
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.TextInput{
						CustomID:    "buyer-location",
						Label:       "Onde você mora?",
						Placeholder: "Pode ser somente a sigla do estado",
						Style:       discordgo.TextInputShort,
						Required:    true,
					},
				},
			},
		},
	}
	response := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &responseData,
	}

	s.InteractionRespond(i.Interaction, &response)
}
