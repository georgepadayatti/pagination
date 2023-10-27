package usecase

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/georgepadayatti/pagination/data_agreement"
	"github.com/georgepadayatti/pagination/data_agreement_record"
	"github.com/georgepadayatti/pagination/db"
	"github.com/georgepadayatti/pagination/policy_author"
	"github.com/georgepadayatti/pagination/revision"
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

func randate() string {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0).UTC().Format("2006-01-02T15:04:05Z")
}

// CreateTenPolicyDocuments Create ten policy documents in the database
func CreateTenPolicyDocuments() {
	for i := 1; i <= 10; i++ {
		// Constructing a policy
		policy := db.Policy{
			Name:      fmt.Sprintf("Policy-%d", i),
			Revisions: createRevisions(i, 10), // creating 10 revisions
			Timestamp: randate(),
		}

		// Creating the policy
		p, err := db.CreatePolicy(policy)
		if err != nil {
			log.Printf("Error creating policy %d: %v", i, err)
			continue // continue to the next iteration
		}
		log.Printf("Policy %d created successfully!", i)

		// Create a policy author object
		CreatePolicyAuthorsDocuments(p.ID.Hex(), i)

	}
}

// CreatePolicyAuthorsDocuments Create policy author documents in the database
func CreatePolicyAuthorsDocuments(policyId string, index int) {

	for i := 1; i <= 5; i++ {
		// Constructing a policy author
		policy := policy_author.PolicyAuthor{
			Timestamp: randate(),
			PolicyId:  policyId,
		}

		// Creating the policy author
		_, err := policy_author.CreatePolicyAuthor(policy)
		if err != nil {
			log.Printf("Error creating policy author %d for policy %d: %v", i, index, err)
		}

		log.Printf("Policy author %d for policy %d created successfully!", i, index)

	}

}

func randNumber(min, max int) int {
	return rand.Intn(max-min) + min
}

// CreateDataAgreement Create data agreement in the database
func CreateDataAgreement() {

	lawfulBases := []string{"consent", "contractual", "legitimate_interest", "public_interest", "vital_task"}

	for i := 1; i <= 5; i++ {
		// Constructing a data agreement
		dataAgreement := data_agreement.DataAgreement{
			Purpose:     fmt.Sprintf("DataAgreement-%d", i),
			LawfulBasis: lawfulBases[randNumber(0, 4)],
		}

		// Creating the data agreement
		d, err := data_agreement.CreateDataAgreement(dataAgreement)
		if err != nil {
			log.Printf("Error creating data agreement %d: %v", i, err)
		}

		log.Printf("Data agreement %d created successfully!", i)

		// Create data agreement records
		CreateDataAgreementRecord(d.Id.Hex(), i)

	}

}

// CreateRevision Create revision in the database
func CreateRevision(dataAgreementRecordId string) {
	for i := 1; i <= 5; i++ {
		// Constructing a revision
		revisionObj := revision.Revision{
			ObjectId:  dataAgreementRecordId,
			Timestamp: randate(),
		}

		// Creating the revision
		_, err := revision.CreateRevision(revisionObj)
		if err != nil {
			log.Printf("Error creating revision %d for data agreement record %v: %v", i, dataAgreementRecordId, err)
		}

		log.Printf("Revision %d for data agreement record %v created successfully!", i, dataAgreementRecordId)

	}
}

// CreateDataAgreementRecord Create data agreement record in the database
func CreateDataAgreementRecord(dataAgreementId string, index int) {

	optIn := []bool{true, false}

	// Constructing a data agreement record
	record := data_agreement_record.DataAgreementRecord{
		DataAgreementId: dataAgreementId,
		IndividualId:    fmt.Sprintf("Individual-%d", index),
		OptIn:           optIn[randNumber(0, 1)],
		OrganisationId:  "Organisation-1",
		IsDeleted:       false,
	}

	// Creating the data agreement record
	r, err := data_agreement_record.CreateDataAgreementRecord(record)
	if err != nil {
		log.Printf("Error creating data agreement record: %v", err)
	}

	log.Printf("Data agreement record created successfully!")

	// Create data agreement record revisions
	CreateRevision(r.Id.Hex())

}
