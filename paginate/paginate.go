package paginate

import (
	"context"
	"errors"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Pagination
type Pagination struct {
	CurrentPage int  `json:"currentPage"`
	TotalItems  int  `json:"totalItems"`
	TotalPages  int  `json:"totalPages"`
	Limit       int  `json:"limit"`
	HasPrevious bool `json:"hasPrevious"`
	HasNext     bool `json:"hasNext"`
}

// PaginationQuery
type PaginationQuery struct {
	Filter      bson.M
	Collection  *mongo.Collection
	Context     context.Context
	CurrentPage int
	Limit       int
}

// PaginatedResult
type PaginatedResult struct {
	Items      interface{} `json:"items"`
	Pagination Pagination  `json:"pagination"`
}

// Paginate
func Paginate(query PaginationQuery, resultSlice interface{}) (*PaginatedResult, error) {

	// Calculate total items
	totalItems, err := query.Collection.CountDocuments(query.Context, query.Filter)
	if err != nil {
		return nil, err
	}

	// Initialize pagination structure
	pagination := Pagination{
		CurrentPage: query.CurrentPage,
		TotalItems:  int(totalItems),
		Limit:       query.Limit,
	}

	// Calculate total pages
	pagination.TotalPages = int(totalItems) / query.Limit
	if int(totalItems)%query.Limit > 0 {
		pagination.TotalPages++
	}

	// Set HasNext and HasPrevious
	pagination.HasPrevious = query.CurrentPage > 1
	pagination.HasNext = query.CurrentPage < pagination.TotalPages

	// Query the database
	opts := options.Find().SetSkip(int64((query.CurrentPage - 1) * query.Limit)).SetLimit(int64(query.Limit))
	cursor, err := query.Collection.Find(query.Context, query.Filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(query.Context)

	// Decode items
	sliceValue := reflect.ValueOf(resultSlice)
	if sliceValue.Kind() != reflect.Ptr || sliceValue.Elem().Kind() != reflect.Slice {
		return nil, errors.New("resultSlice must be a slice pointer")
	}
	sliceElem := sliceValue.Elem()
	itemTyp := sliceElem.Type().Elem()

	for cursor.Next(query.Context) {
		itemPtr := reflect.New(itemTyp).Interface()
		if err := cursor.Decode(itemPtr); err != nil {
			return nil, err
		}
		sliceElem = reflect.Append(sliceElem, reflect.ValueOf(itemPtr).Elem())
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &PaginatedResult{
		Items:      sliceElem.Interface(),
		Pagination: pagination,
	}, nil
}
