package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	ServerNumber int                `bson:"server" json:"server" validate:"required"`
	Maintenance  bool               `bson:"maintainance" json:"maintainance" default:"false"`
	APIKey       string             `bson:"api_key,omitempty" json:"api_key"`
	Token        string             `bson:"token,omitempty" json:"token"`
	ExchangeRate float64            `bson:"exchangeRate,omitempty" json:"exchangeRate" default:"0.0"`
	Margin       float64            `bson:"margin,omitempty" json:"margin" default:"0.0"`
	CreatedAt    time.Time          `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt,omitempty" json:"updatedAt"`
}

func InitializeServerCollection(db *mongo.Database) *mongo.Collection {
	collection := db.Collection("servers")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = EnsureIndexes(ctx, collection)

	return collection
}

func EnsureIndexes(ctx context.Context, collection *mongo.Collection) error {
	indexModel := mongo.IndexModel{
		Keys:    map[string]interface{}{"userId": 1},
		Options: options.Index().SetUnique(true),
	}
	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	return err
}
