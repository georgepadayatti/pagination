package data_agreement_record

import (
	"context"

	"github.com/georgepadayatti/pagination/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DataAgreementRecord struct {
	Id              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	DataAgreementId string             `json:"dataAgreementId"`
	IndividualId    string             `json:"individualId"`
	OptIn           bool               `json:"optIn"`
	OrganisationId  string             `json:"-"`
	IsDeleted       bool               `json:"-"`
}

// Collection Gets the handles for collection
func Collection() *mongo.Collection {
	return db.DB.Client.Database(db.DB.Name).Collection("dataAgreementRecords1")
}

// CreateDataAgreementRecord
func CreateDataAgreementRecord(dataAgreementRecord DataAgreementRecord) (DataAgreementRecord, error) {
	dataAgreementRecord.Id = primitive.NewObjectID()
	_, err := Collection().InsertOne(context.TODO(), &dataAgreementRecord)
	if err != nil {
		return dataAgreementRecord, err
	}
	return dataAgreementRecord, nil
}

// CreatePipelineForFilteringDataAgreementRecords This pipeline is used for filtering data agreement records by `id` and `lawfulBasis`
// `id` has 3 possible values - dataAgreementRecordId, dataAgreementId, individualId
func CreatePipelineForFilteringDataAgreementRecords(organisationId string, id string, lawfulBasis string) ([]primitive.M, error) {

	var pipeline []bson.M

	// Stage 1 - Match by `organisationId` and `isDeleted=false`
	pipeline = append(pipeline, bson.M{"$match": bson.M{"organisationid": organisationId, "isdeleted": false}})

	if len(id) > 0 {

		or := []bson.M{
			{"dataagreementid": id},
			{"individualid": id},
		}

		// Stage 2 - Match `id` against `dataAgreementRecordId`, `dataAgreementId`, `individualId`
		convertIdtoObjectId, err := primitive.ObjectIDFromHex(id)
		if err == nil {
			// Append `dataAgreementRecordId` `or` statements only if
			// string is converted to objectId without errors
			or = append(or, bson.M{"_id": convertIdtoObjectId})
		}

		pipeline = append(pipeline, bson.M{"$match": bson.M{
			"$or": or,
		}})
	}

	// Stage 3 - Lookup data agreement document by `dataAgreementId`
	// This is done to obtain `policy` and `lawfulBasis` fields from data agreement document
	pipeline = append(pipeline, bson.M{"$lookup": bson.M{
		"from": "dataAgreements1",
		"let":  bson.M{"localId": "$dataagreementid"},
		"pipeline": bson.A{
			bson.M{
				"$match": bson.M{
					"$expr": bson.M{
						"$eq": []interface{}{"$_id", bson.M{"$toObjectId": "$$localId"}},
					},
				},
			},
		},
		"as": "dataAgreements",
	}})

	// Stage 4 - Unwind the data agreement fields
	pipeline = append(pipeline, bson.M{"$unwind": "$dataAgreements"})

	// Stage 5 - Lookup revision by `dataAgreementRecordId`
	// This is done to obtain timestamp for the latest revision of the data agreement record.
	pipeline = append(pipeline, bson.M{"$lookup": bson.M{
		"from": "revisions1",
		"let":  bson.M{"localId": "$_id"},
		"pipeline": bson.A{
			bson.M{
				"$match": bson.M{
					"$expr": bson.M{
						"$eq": []interface{}{"$objectid", bson.M{"$toString": "$$localId"}},
					},
				},
			},
			bson.M{
				"$sort": bson.M{"timestamp": -1},
			},
			bson.M{"$limit": int64(1)},
		},
		"as": "revisions",
	}})

	// Stage 6 - Add the timestamp from revisions
	pipeline = append(pipeline, bson.M{"$addFields": bson.M{"timestamp": bson.M{
		"$let": bson.M{
			"vars": bson.M{
				"first": bson.M{
					"$arrayElemAt": bson.A{"$revisions", 0},
				},
			},
			"in": "$$first.timestamp",
		},
	}}})

	// Stage 7 - Remove revisions field
	pipeline = append(pipeline, bson.M{
		"$project": bson.M{
			"revisions": 0,
		},
	})

	// Stage 8 - Match by lawful basis
	if len(lawfulBasis) > 0 {
		pipeline = append(pipeline, bson.M{
			"$match": bson.M{
				"dataAgreements.lawfulbasis": lawfulBasis,
			},
		})
	}

	return pipeline, nil
}
