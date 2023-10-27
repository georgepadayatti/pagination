package data_agreement

import (
	"context"

	"github.com/georgepadayatti/pagination/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DataAgreement struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Purpose     string             `json:"purpose" valid:"required"`
	LawfulBasis string             `json:"lawfulBasis" valid:"required"`
}

// Collection Gets the handles for collection
func Collection() *mongo.Collection {
	return db.DB.Client.Database(db.DB.Name).Collection("dataAgreements1")
}

// CreateDataAgreement
func CreateDataAgreement(dataAgreement DataAgreement) (DataAgreement, error) {
	dataAgreement.Id = primitive.NewObjectID()
	_, err := Collection().InsertOne(context.TODO(), &dataAgreement)
	if err != nil {
		return dataAgreement, err
	}
	return dataAgreement, nil
}
