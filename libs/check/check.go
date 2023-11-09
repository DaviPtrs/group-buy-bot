package check

import (
	"fmt"
	"strings"

	"github.com/DaviPtrs/group-buy-bot/libs/approval"
	"github.com/DaviPtrs/group-buy-bot/libs/bot/session"
	"github.com/DaviPtrs/group-buy-bot/libs/mongorm"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var StateColletionName = "states"

func init() {
	client := mongorm.ConnectedClient()
	defer mongorm.DisconnectClient(client)
	coll := client.Database(mongorm.DatabaseName).Collection(StateColletionName)

	trueVar := true
	indexes := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "guild_id", Value: 1}},
			Options: &options.IndexOptions{Unique: &trueVar},
		},
	}
	mongorm.AddIndexes(coll, indexes)

	var state *GBState = new(GBState)
	err := state.Read(coll, stateFilter, state)

	if err == mongo.ErrNoDocuments {
		logrus.Info("No state found on DB. Populating...")
		state.PopulateState()
		err := state.Create(coll, state)
		if err != nil {
			logrus.Errorf("Unable to create state: %v", err)
		} else {
			logrus.Info("State initialized")
		}
	}

	if state.GuildID != session.GetGuildID() {
		logrus.Fatalf("Stored state is invalid! %v", *state)
	}
}

func GetDistinctBuyers() []string {
	client := mongorm.ConnectedClient()
	defer mongorm.DisconnectClient(client)

	coll := client.Database(mongorm.DatabaseName).Collection(approval.ApprovedCollectionName)

	results, err := coll.Distinct(mongorm.DefaultContext, "item.user_id", bson.D{})
	if err != nil {
		logrus.Errorf("Failed to fetch distinct \"item.user_id\": %v", err)
		return nil
	}

	var res []string
	for _, x := range results {
		value, ok := x.(string)
		if !ok {
			logrus.Errorf("Unable to convert value %v to string", ok)
			return nil
		}
		res = append(res, value)
	}

	return res
}

func CheckReadyToBuy() {
	buyers := GetDistinctBuyers()
	if buyers == nil {
		logrus.Errorf("buyers list on checkReady is null")
		return
	}

	state := GetState()
	count := len(buyers)

	if count < 2 {
		state.ReadyMessageSent = false
		state.Save()
		return
	}

	if count != state.LastBuyerCount {
		state.ReadyMessageSent = false
		state.LastBuyerCount = count
		state.Save()
	}

	if state.ReadyMessageSent {
		return
	}

	var sb strings.Builder
	sb.WriteString("**ATENÇÃO: ** O Group buy atingiu 2 compradores ou mais\n")
	sb.WriteString("Os compradores listados abaixo podem combinar quando efetuarão as compras.\n")
	sb.WriteString("Lembrem de ler o regulamento do Group Buy e não esqueçam de enviar o código de rastreio pros admins.\n")
	sb.WriteString("\n**COMPRADORES**\n")

	for _, buyer := range buyers {
		sb.WriteString(fmt.Sprintf("<@%s>\n", buyer))
	}

	s := session.GetDiscordSession()
	for _, buyer := range buyers {
		st, err := s.UserChannelCreate(buyer)
		if err != nil {
			logrus.Errorf("failed to create DM with user %v: %v", buyer, err)
			return
		}

		_, err = s.ChannelMessageSend(st.ID, sb.String())
		if err != nil {
			logrus.Errorf("failed to send gb_ready message to user %v: %v", buyer, err)
			return
		}
	}
	state.ReadyMessageSent = true
	state.Save()
}
