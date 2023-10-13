package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/georgepadayatti/pagination/db"
	"github.com/georgepadayatti/pagination/usecase"
)

func main() {
	err := db.Init()
	if err != nil {
		panic(err)
	}
	log.Println("Database session opened")

	// usecase.CreateTenPolicyDocuments()

	result, err := usecase.GetPaginatedRevisionsFromDb()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Convert struct to JSON with indents.
	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	// Convert JSON bytes to string and print.
	fmt.Println(string(jsonData))
}
