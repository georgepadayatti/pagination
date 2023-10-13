package usecase

import (
	"context"

	"github.com/georgepadayatti/pagination/db"
	"github.com/georgepadayatti/pagination/paginate"
	"go.mongodb.org/mongo-driver/bson"
)

func GetPaginatedRevisionsFromDb() (*paginate.PaginatedResult, error) {
	query := paginate.PaginationQuery{
		Filter:      bson.M{},
		Collection:  db.Collection(),
		Context:     context.Background(),
		CurrentPage: 1,
		Limit:       5,
	}
	var policies []db.Policy
	result, err := paginate.Paginate(query, &policies)

	return result, err
}
