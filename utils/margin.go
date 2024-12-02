package utils

import (
	"context"
	"fmt"

	"github.com/ranjankuldeep/updateServicesList/internal/database/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FetchMarginAndExchangeRate(ctx context.Context, db *mongo.Database) (map[int]float64, map[int]float64, error) {
	serverCollection := models.InitializeServerCollection(db)
	marginMap := make(map[int]float64)
	exchangeRateMap := make(map[int]float64)

	cursor, err := serverCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch servers: %w", err)
	}
	defer cursor.Close(ctx)

	// Iterate over the fetched servers and populate the maps
	for cursor.Next(ctx) {
		var server models.Server
		if err := cursor.Decode(&server); err != nil {
			return nil, nil, fmt.Errorf("failed to decode server: %w", err)
		}
		marginMap[server.ServerNumber] = server.Margin
		exchangeRateMap[server.ServerNumber] = server.ExchangeRate
	}

	if err := cursor.Err(); err != nil {
		return nil, nil, fmt.Errorf("error while iterating over servers: %w", err)
	}
	return marginMap, exchangeRateMap, nil
}
