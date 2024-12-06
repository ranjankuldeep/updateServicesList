package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/ranjankuldeep/updateServicesList/internal/database"
	"github.com/ranjankuldeep/updateServicesList/internal/runner"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")
	databaseName := os.Getenv("MONGODB_DATABASE")
	uri := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.blfflhg.mongodb.net/%s?retryWrites=true&w=majority", username, password, databaseName)

	client, err := database.ConnectDB(uri)
	if err != nil {
		log.Fatal("Error initializing MongoDB connection:", err)
	}
	db := client.Database(databaseName)
	startTime := time.Now()
	runner.UpdateServerData(db, context.TODO())

	duration := time.Since(startTime).Seconds()
	durationRoundOff := fmt.Sprintf("%.2f seconds", duration)
	fmt.Println("Time taken to update server data:", durationRoundOff)
}
