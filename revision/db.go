package revision

import (
	"context"

	"github.com/georgepadayatti/pagination/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Revision
type Revision struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ObjectId  string             `json:"objectId"`
	Timestamp string             `json:"timestamp"`
}

// Collection Gets the handles for collection
func Collection() *mongo.Collection {
	return db.DB.Client.Database(db.DB.Name).Collection("revisions1")
}

// CreateRevision
func CreateRevision(revision Revision) (Revision, error) {
	revision.Id = primitive.NewObjectID()
	_, err := Collection().InsertOne(context.TODO(), &revision)
	if err != nil {
		return revision, err
	}
	return revision, nil
}
