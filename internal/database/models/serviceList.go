package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServerData struct {
	Server int    `bson:"server" json:"server"`
	Price  string `bson:"price" json:"price"`
	Code   string `bson:"code" json:"code"`
	Otp    string `bson:"otp" json:"otp"`
	Block  bool   `bson:"block" json:"block" default:"false"`
}

type ServerList struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name" json:"name"`
	Service_Code string             `bson:"service_code" json:"service_code"`
	Servers      []ServerData       `bson:"servers" json:"servers"`
	CreatedAt    time.Time          `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt,omitempty" json:"updatedAt"`
}

func InitializeServerListCollection(db *mongo.Database) *mongo.Collection {
	return db.Collection("serverlists")
}
