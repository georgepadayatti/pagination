package usecase

import (
	"fmt"
	"log"

	"github.com/georgepadayatti/pagination/db"
)

// createRevisions creates a specified number (count) of Revision structs,
// starting their IDs and Snapshots with a base number (baseID).
func createRevisions(baseID, count int) []db.Revision {
	revisions := make([]db.Revision, count)
	for j := 0; j < count; j++ {
		revisions[j] = db.Revision{
			ID:                 fmt.Sprintf("RevID-%d-%d", baseID, j),
			SerialisedSnapshot: fmt.Sprintf("SnapshotData-%d-%d", baseID, j),
		}
	}
	return revisions
}

// CreateTenPolicyDocuments Create ten policy documents in the database
func CreateTenPolicyDocuments() {
	for i := 1; i <= 10; i++ {
		// Constructing a policy
		policy := db.Policy{
			Name:      fmt.Sprintf("Policy-%d", i),
			Revisions: createRevisions(i, 10), // creating 10 revisions
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
