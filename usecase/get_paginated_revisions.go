package usecase

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/georgepadayatti/pagination/db"
	"github.com/georgepadayatti/pagination/paginate"
)

func GetPaginatedRevisions() {
	policy, err := db.ReadFirstPolicy()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	toBeSortedItems := make([]interface{}, len(policy.Revisions))
	for i, r := range policy.Revisions {
		toBeSortedItems[i] = r
	}

	query := paginate.PaginateObjectsQuery{
		Limit:  2,
		Offset: 0,
	}
	result := paginate.PaginateObjects(query, toBeSortedItems)

	// Convert struct to JSON with indents.
	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	// Convert JSON bytes to string and print.
	fmt.Println(string(jsonData))
}
