package item

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Item struct {
	CustomID      string  `bson:"custom_id"`
	UserID        string  `bson:"user_id"`
	URL           string  `bson:"url"`
	Price         float32 `bson:"price"`
	Weight        float32 `bson:"weight"`
	TaxRate       int     `bson:"tax_rate"`
	BuyerLocation string  `bson:"buyer_location"`
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

	userID := UserIDfromItemCustomID(data.CustomID)

	return &Item{data.CustomID, userID, url.String(), float32(price), float32(weight), int(tax), location}, nil
}

func (item *Item) ParseToEmbedFields() *[]*discordgo.MessageEmbedField {
	fields := []*discordgo.MessageEmbedField{
		{
			Name: "ID", Value: item.CustomID,
		},
		{
			Name: "URL", Value: item.URL,
		},
		{
			Name: "Price", Value: fmt.Sprintf("$%.2f", item.Price),
		},
		{
			Name: "Weight", Value: fmt.Sprintf("%.2flbs", item.Weight),
		},
		{
			Name: "Estimated Tax", Value: fmt.Sprintf("%d%%", item.TaxRate),
		},
		{
			Name: "Buyer's location", Value: item.BuyerLocation,
		},
	}

	return &fields
}

func UserIDfromItemCustomID(customID string) string {

	re := regexp.MustCompile(`group_buy_item_([0-9]+)_`)

	matches := re.FindStringSubmatch(customID)

	if len(matches) == 2 {
		return matches[1]
	}
	return ""
}

func (i *Item) GetModel() *ItemModel {
	var model *ItemModel = new(ItemModel)
	model.CustomID = i.CustomID
	model.UserID = i.UserID
	model.URL = i.URL
	model.Price = i.Price
	model.Weight = i.Weight
	model.TaxRate = i.TaxRate
	model.BuyerLocation = i.BuyerLocation
	return model
}
