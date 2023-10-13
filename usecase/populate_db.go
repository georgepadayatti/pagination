package usecase

import (
	"fmt"
	"log"

	"github.com/georgepadayatti/pagination/db"
)

// CreateTenPolicyDocuments Create ten policy documents in the database
func CreateTenPolicyDocuments() {
	for i := 1; i <= 10; i++ {
		// Constructing a policy
		policy := db.Policy{
			Name: fmt.Sprintf("Policy-%d", i),
			Revisions: []db.Revision{
				{
					ID:                 fmt.Sprintf("RevID-%d", i),
					SerialisedSnapshot: fmt.Sprintf("SnapshotData-%d", i),
				},
			},
		}

		// Creating the policy
		_, err := db.CreatePolicy(policy)
		if err != nil {
			log.Printf("Error creating policy %d: %v", i, err)
			continue // continue to the next iteration
		}

		log.Printf("Policy %d created successfully!", i)
	}
}
