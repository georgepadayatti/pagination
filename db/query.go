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
	Timestamp string             `json:"timestamp"`
}

// Collection Gets the handles for collection
func Collection() *mongo.Collection {
	return DB.Client.Database(DB.Name).Collection(config.AppConfig.GetPolicyCollectionName())
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

// QueryPoliciesByName
func QueryPoliciesByName(name string) ([]Policy, error) {

	// match := bson.M{"$match": bson.M{"name": name}}
	sort := bson.M{"$sort": bson.M{"timestamp": -1}}
	pipeline := []bson.M{sort}

	var policies []Policy

	cursor, err := Collection().Aggregate(context.TODO(), pipeline)
	if err != nil {
		return policies, err
	}
	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &policies); err != nil {
		return policies, err
	}
	return policies, nil
}
