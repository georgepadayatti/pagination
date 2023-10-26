package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/georgepadayatti/pagination/paginate"
	"github.com/georgepadayatti/pagination/policy_author"
	"go.mongodb.org/mongo-driver/bson"
)

func GetPoliciesByName() {
	const policyId = "653a8d291ae8ac60a966cb1b"

	pipeline := []bson.M{
		{"$match": bson.M{"policyid": policyId}},
		{"$lookup": bson.M{
			"from": "dummyCollection",
			"let": bson.M{
				"localId": "$_id",
			},
			"pipeline": bson.A{
				bson.M{
					"$match": bson.M{
						"$expr": bson.M{
							"$eq": []interface{}{"$policyid", "$$localId"},
						},
					},
				},
			},
			"as": "policy",
		}},
		{
			"$sort": bson.M{"timestamp": -1},
		},
	}

	query := paginate.PaginateDBObjectsQuery{
		Pipeline:   pipeline,
		Collection: policy_author.Collection(),
		Context:    context.Background(),
		Limit:      5,
		Offset:     0,
	}
	var policies []policy_author.PolicyAuthor
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
