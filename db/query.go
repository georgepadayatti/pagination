package db

import (
	"context"

	"github.com/georgepadayatti/pagination/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Revision
type Revision struct {
	ID                 string `json:"id"`
	SerialisedSnapshot string `json:"serialisedSnapshot"`
}

// Policy
type Policy struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name"`
	Revisions []Revision         `json:"revisions"`
}

// Collection Gets the handles for collection
func Collection() *mongo.Collection {
	return DB.Client.Database(DB.Name).Collection(config.AppConfig.GetCollectionName())
}

// CreatePolicy Create policy
func CreatePolicy(policy Policy) (Policy, error) {
	policy.ID = primitive.NewObjectID()
	_, err := Collection().InsertOne(context.TODO(), &policy)
	if err != nil {
		return policy, err
	}
	return policy, nil
}

// ReadFirstPolicy Read first policy
func ReadFirstPolicy() (Policy, error) {

	var policy Policy
	err := Collection().FindOne(context.TODO(), bson.M{}).Decode(&policy)

	return policy, err
}
