package admin

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"

	"github.com/DaviPtrs/group-buy-bot/libs/approval"
	"github.com/DaviPtrs/group-buy-bot/libs/item"
	"github.com/DaviPtrs/group-buy-bot/libs/mongorm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FetchAllItems() *[]*item.ItemModel {
	client := mongorm.ConnectedClient()
	defer mongorm.DisconnectClient(client)

	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{Key: "item.user_id", Value: 1}})

	coll := client.Database(mongorm.DatabaseName).Collection(approval.ApprovedCollectionName)
	cursor, err := coll.Find(mongorm.DefaultContext, filter, opts)
	if err != nil {
		log.Panic(err)
	}

	var results []*item.ItemModel
	if err = cursor.All(mongorm.DefaultContext, &results); err != nil {
		log.Panic(err)
	}

	return &results
}

func GenerateCSV(modelList *[]*item.ItemModel) []byte {
	headers := []string{
		"User",
		"Price",
		"Weight (lbs)",
		"Tax rate",
		"Buyer Location",
		"URL",
	}

	var rows [][]string

	for _, i := range *modelList {
		row := []string{
			i.UserID,
			fmt.Sprintf("%.2f", i.Price),
			fmt.Sprintf("%.2f", i.Weight),
			fmt.Sprintf("%.2f", float32(i.TaxRate)/100),
			i.BuyerLocation,
			i.URL,
		}
		rows = append(rows, row)
	}

	var buffer bytes.Buffer
	w := csv.NewWriter(&buffer)

	w.Write(headers)
	w.WriteAll(rows)

	return buffer.Bytes()
}
