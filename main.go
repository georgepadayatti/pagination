package main

import (
	"fmt"
	"log"

	"github.com/georgepadayatti/pagination/cmd"
	"github.com/georgepadayatti/pagination/config"
	"github.com/georgepadayatti/pagination/db"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	dbUser := config.AppConfig.GetDatabaseUser()
	dbPassword := config.AppConfig.GetDatabasePassword()
	dbName := config.AppConfig.GetDatabaseName()
	collectionName := config.AppConfig.GetCollectionName()
	err = db.Init(dbUser, dbPassword, dbName, collectionName)
	if err != nil {
		panic(err)
	}
	log.Println("Database session opened")

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}

}
