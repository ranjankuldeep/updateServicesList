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
	databaseName := os.Getenv("MONGODB_DATABASE")
	uri := os.Getenv("MONGODB_URI")
	client, err := database.ConnectDB(databaseName, uri)
	if err != nil {
		log.Fatal("Error initializing MongoDB connection:", err)
	}
	db := client.Database(databaseName)

	startTime := time.Now()
	runner.UpdateServerData(db, context.TODO())
	duration := time.Since(startTime).Seconds()

	durationRoundOff := fmt.Sprintf("%.2f seconds", duration)
	fmt.Println("Time taken to update server data:", durationRoundOff)
	fmt.Println("Succesfully updated servicelist data")
}
