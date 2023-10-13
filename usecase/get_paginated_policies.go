package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/georgepadayatti/pagination/db"
	"github.com/georgepadayatti/pagination/paginate"
	"go.mongodb.org/mongo-driver/bson"
)

func GetPaginatedPolicies() {
	query := paginate.PaginateDBObjectsQuery{
		Filter:     bson.M{},
		Collection: db.Collection(),
		Context:    context.Background(),
		Limit:      1,
		Offset:     0,
	}
	var policies []db.Policy
	result, err := paginate.PaginateDBObjects(query, &policies)
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
