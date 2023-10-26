package policy_author

import (
	"context"

	"github.com/georgepadayatti/pagination/config"
	"github.com/georgepadayatti/pagination/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// PolicyAuthor
type PolicyAuthor struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Timestamp string             `json:"timestamp"`
	PolicyId  string             `json:"policyId"`
}

// Collection Gets the handles for collection
func Collection() *mongo.Collection {
	return db.DB.Client.Database(db.DB.Name).Collection(config.AppConfig.GetPolicyAuthorCollectionName())
}

// CreatePolicyAuthor Create policy author
func CreatePolicyAuthor(policyAuthor PolicyAuthor) (PolicyAuthor, error) {
	policyAuthor.ID = primitive.NewObjectID()
	_, err := Collection().InsertOne(context.TODO(), &policyAuthor)
	if err != nil {
		return policyAuthor, err
	}
	return policyAuthor, nil
}

// ReadFirstPolicy Read first policy
func ReadFirstPolicy() (PolicyAuthor, error) {

	var policyAuthor PolicyAuthor
	err := Collection().FindOne(context.TODO(), bson.M{}).Decode(&policyAuthor)

	return policyAuthor, err
}
