package user

import (
	"fmt"
	"strings"

	"github.com/DaviPtrs/group-buy-bot/libs/approval"
	"github.com/DaviPtrs/group-buy-bot/libs/bot/session"
	"github.com/DaviPtrs/group-buy-bot/libs/mongorm"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

var readyMessageSent = false
var lastBuyerCount = 0

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
	count := len(buyers)
	if count < 2 {
		readyMessageSent = false
		return
	}

	if count != lastBuyerCount {
		readyMessageSent = false
		lastBuyerCount = count
	}

	if readyMessageSent {
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
	readyMessageSent = true
}
