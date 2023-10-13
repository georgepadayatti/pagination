package db

import (
	"context"

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
	return DB.Client.Database(DB.Name).Collection(COLLECTION_NAME)
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
