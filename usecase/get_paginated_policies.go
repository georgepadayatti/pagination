package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/georgepadayatti/pagination/db"
	"github.com/georgepadayatti/pagination/paginate"
)

func GetPaginatedPolicies() {
	query := paginate.PaginateDBObjectsQuery{
		Collection: db.Collection(),
		Context:    context.Background(),
		Limit:      1,
		Offset:     1,
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
