package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/georgepadayatti/pagination/data_agreement_record"
	"github.com/georgepadayatti/pagination/paginate"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DataAgreementForListDataAgreementRecord struct {
	Purpose     string `json:"purpose"`
	LawfulBasis string `json:"lawfulBasis"`
}

type DataAgreementRecordForAuditList struct {
	Id              primitive.ObjectID                      `json:"id" bson:"_id,omitempty"`
	DataAgreementId string                                  `json:"dataAgreementId"`
	IndividualId    string                                  `json:"individualId"`
	OptIn           bool                                    `json:"optIn"`
	DataAgreements  DataAgreementForListDataAgreementRecord `json:"dataAgreement"`
	Timestamp       string                                  `json:"timestamp"`
}

func GetDataAgreementRecords() {
	const id = "653bbda3b48db16a5b7c2a3d"
	const lawfulBasis = "consent"

	pipeline, err := data_agreement_record.CreatePipelineForFilteringDataAgreementRecords("Organisation-1", id, lawfulBasis)
	if err != nil {
		log.Fatalf("Error while creating pipeline: %v", err)
	}

	query := paginate.PaginateDBObjectsQuery{
		Pipeline:   pipeline,
		Collection: data_agreement_record.Collection(),
		Context:    context.Background(),
		Limit:      5,
		Offset:     0,
	}
	var res []DataAgreementRecordForAuditList
	result, err := paginate.PaginateDBObjects(query, &res)
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
