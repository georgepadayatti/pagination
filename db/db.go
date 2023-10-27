package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type db struct {
	Client *mongo.Client
	Name   string
}

// DB Database session pointer
var DB db

// Init Initialises the database connection
func Init(dbUser string, dbPassword string, dbName string, collectionName ...string) error {
	dbUrl := fmt.Sprintf("mongodb://%s:%s@localhost:27017/%s", dbUser, dbPassword, dbName)
	clientOptions := options.Client().ApplyURI(dbUrl)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a new MongoDB client
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Printf("Error connecting to MongoDB: %v", err)
		return err
	}

	// Ping the MongoDB server
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	DB = db{
		Client: client,
		Name:   dbName,
	}

	for i := 0; i < len(collectionName); i++ {
		err = initCollection(collectionName[i], []string{"name"}, true)
		if err != nil {
			log.Printf("Initialising collection: %v", err)
			return err
		}
	}

	err = initCollection("policyAuthors", []string{"name"}, true)
	if err != nil {
		log.Printf("Initialising collection: %v", err)
		return err
	}

	err = initCollection("revisions1", []string{"name"}, true)
	if err != nil {
		log.Printf("Initialising collection: %v", err)
		return err
	}

	err = initCollection("dataAgreements1", []string{"name"}, true)
	if err != nil {
		log.Printf("Initialising collection: %v", err)
		return err
	}

	err = initCollection("dataAgreementRecords1", []string{"name"}, true)
	if err != nil {
		log.Printf("Initialising collection: %v", err)
		return err
	}

	return nil
}

// Init Initialises collection
func initCollection(collectionName string, keys []string, unique bool) error {

	c := DB.Client.Database(DB.Name).Collection(collectionName)

	indexOptions := options.Index()

	keysDoc := bson.D{}
	for _, key := range keys {
		keysDoc = append(keysDoc, bson.E{Key: key, Value: 1})
	}

	indexModel := mongo.IndexModel{
		Keys:    keysDoc,
		Options: indexOptions.SetSparse(true).SetUnique(unique),
	}

	_, err := c.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		log.Printf("Error creating index on the specified keys: %v", err)
		return err
	}

	log.Printf("Initialized %v collection", collectionName)
	return nil
}
