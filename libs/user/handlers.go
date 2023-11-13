package user

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/DaviPtrs/group-buy-bot/libs/approval"
	"github.com/DaviPtrs/group-buy-bot/libs/item"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

var UserChannelID string

func init() {
	godotenv.Load()

	var ok bool
	UserChannelID, ok = os.LookupEnv("DISCORD_BOT_USER_CHANNEL_ID")
	if !ok {
		logrus.Fatal("User Channel ID not found")
	}
}

func CommandHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"add": addCommandHandler,
	}
}

func ModalHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"add": addModalHandler,
	}
}

func addCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ChannelID != UserChannelID {
		WrongChannelResponse(s, i.Interaction, UserChannelID)
		return
	}

	uuid := uuid.NewV4()
	responseData := discordgo.InteractionResponseData{
		CustomID: "group_buy_item_" + i.Interaction.Member.User.ID + "_" + uuid.String(),
		Title:    "Group buy - Request new product",
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.TextInput{
						CustomID:    "product-url",
						Label:       "Link do produto",
						Style:       discordgo.TextInputShort,
						Placeholder: "https://www.amazon.com/Images-You-Should-Not-Masturbate/dp/0399536493",
						Required:    true,
						MaxLength:   4000,
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
						MaxLength:   5,
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
						MaxLength:   5,
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
						MaxLength:   5,
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
						MaxLength:   50,
					},
				},
			},
		},
	}
	response := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &responseData,
	}

	err := s.InteractionRespond(i.Interaction, &response)
	if err != nil {
		logrus.Errorf("Error on responding with modal: %v", err)
	}
}

func addModalHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ModalSubmitData()
	if !strings.HasPrefix(data.CustomID, "group_buy_item") {
		return
	}

	logrus.Infof("Item received from user %v", i.Member.User.String())
	err := approval.SendItemToApproval(s, i.Member.User.ID, &data)
	var submit_message string
	if err != nil {
		valErr, ok := err.(*item.InvalidItem)
		if ok {
			submit_message = "Não foi possível adicionar seu item na lista.\n"
			submit_message += fmt.Sprintf("Campo \"%v\" é inválido!", valErr.InvalidField)
		} else {
			logrus.Errorf("Unable to send item to approval: %v", err)
			return
		}
	} else {
		submit_message = "Obrigado por enviar seu item pra lista do group buy.\n"
		submit_message += "Os admins irão analisar se seu item é valido e você receberá uma confirmação na DM."
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: submit_message,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		logrus.Errorf("Unable to respond to modal %v: %v", data.CustomID, err)
		return
	}

	time.Sleep(time.Second * 10)
	err = s.InteractionResponseDelete(i.Interaction)
	if err != nil {
		logrus.Errorf("Error on deleting \"%v modal confirmation\" response: %v", data.CustomID, err)
	}
}
