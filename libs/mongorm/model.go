package mongorm

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Model struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func (m *Model) Create(ctx context.Context, coll *mongo.Collection, model interface{}) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	res, err := coll.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	m.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (m *Model) Read(ctx context.Context, coll *mongo.Collection, filter interface{}, result interface{}) error {
	err := coll.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

func (m *Model) Update(ctx context.Context, coll *mongo.Collection, filter interface{}, update interface{}) error {
	m.UpdatedAt = time.Now()

	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (m *Model) Delete(ctx context.Context, coll *mongo.Collection, filter interface{}) error {
	_, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
