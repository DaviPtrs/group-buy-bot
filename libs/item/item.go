package item

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Item struct {
	CustomID      string
	URL           string
	Price         float32
	Weight        float32
	TaxRate       int
	BuyerLocation string
}

func ParseFromModal(data *discordgo.ModalSubmitInteractionData) (*Item, error) {

	url, err := url.ParseRequestURI(data.Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value)
	if err != nil {
		return nil, &InvalidItem{InvalidField: "product-url", Err: err}
	}

	numbersRegex, _ := regexp.Compile(`[-+]?(?:\d*\.*\d+)`)

	pricePlain := data.Components[1].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value
	pricePlain = strings.ReplaceAll(pricePlain, ",", ".")
	priceStr := numbersRegex.FindString(pricePlain)
	price, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		return nil, &InvalidItem{InvalidField: "price", Err: err}
	}

	weightPlain := data.Components[2].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value
	weightPlain = strings.ToLower(weightPlain)
	weightPlain = strings.ReplaceAll(weightPlain, ",", ".")
	weightStr := numbersRegex.FindString(weightPlain)
	weight, err := strconv.ParseFloat(weightStr, 32)
	if err != nil {
		return nil, &InvalidItem{InvalidField: "weight", Err: err}
	}
	if strings.Contains(weightPlain, "k") {
		weight *= 2.205
	}

	taxPlain := data.Components[3].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value
	taxPlain = strings.ToLower(taxPlain)
	tax, err := strconv.ParseInt(numbersRegex.FindString(taxPlain), 10, 0)
	if err != nil {
		return nil, &InvalidItem{InvalidField: "tax-rate", Err: err}
	}

	location := data.Components[4].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value

	return &Item{data.CustomID, url.String(), float32(price), float32(weight), int(tax), location}, nil
}
