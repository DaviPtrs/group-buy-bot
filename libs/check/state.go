package check

import (
	"github.com/DaviPtrs/group-buy-bot/libs/bot/session"
	"github.com/DaviPtrs/group-buy-bot/libs/mongorm"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var stateFilter = bson.M{"guild_id": session.GetGuildID()}

type GBState struct {
	mongorm.Model
	GuildID          string `bson:"guild_id"`
	ReadyMessageSent bool   `bson:"ready_message_sent"`
	LastBuyerCount   int    `bson:"last_buyer_count"`
}

func (s *GBState) PopulateState() {
	s.GuildID = session.GetGuildID()
	s.ReadyMessageSent = false
	s.LastBuyerCount = 0
}

func (s *GBState) Save() {
	client := mongorm.ConnectedClient()
	defer mongorm.DisconnectClient(client)
	coll := client.Database(mongorm.DatabaseName).Collection(StateColletionName)

	update := bson.M{
		"$set": bson.M{
			"ready_message_sent": s.ReadyMessageSent,
			"last_buyer_count":   s.LastBuyerCount,
			"updated_at":         primitive.NewDateTimeFromTime(s.UpdatedAt),
		},
	}
	err := s.Update(coll, stateFilter, update)

	if err != nil {
		logrus.Fatalf("Unable to save state to DB: %v", err)
	}
}

func GetState() *GBState {
	client := mongorm.ConnectedClient()
	defer mongorm.DisconnectClient(client)
	coll := client.Database(mongorm.DatabaseName).Collection(StateColletionName)

	var state *GBState = new(GBState)
	err := state.Read(coll, stateFilter, state)

	if err != nil {
		logrus.Fatalf("Unable to fetch state from DB: %v", err)
	}

	return state
}
