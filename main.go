package main

import (
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
	// usecase.GetPaginatedPolicies()
	usecase.GetPaginatedRevisions()

}
