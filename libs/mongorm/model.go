package mongorm

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var DefaultContext = context.Background()

type Model struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func (m *Model) Create(coll *mongo.Collection, model interface{}) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	res, err := coll.InsertOne(DefaultContext, model)
	if err != nil {
		return err
	}

	m.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (m *Model) Read(coll *mongo.Collection, filter interface{}, result interface{}) error {
	err := coll.FindOne(DefaultContext, filter).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

func (m *Model) Update(coll *mongo.Collection, filter interface{}, update interface{}) error {
	m.UpdatedAt = time.Now()

	_, err := coll.UpdateOne(DefaultContext, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (m *Model) Delete(coll *mongo.Collection, filter interface{}) error {
	_, err := coll.DeleteOne(DefaultContext, filter)
	if err != nil {
		return err
	}

	return nil
}
