package db

import (
	"context"

	"github.com/rohan-luthra/microservice-grpc-go/service-restaurants-go/src/kafka"
	"github.com/rohan-luthra/microservice-grpc-go/service-restaurants-go/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db *Database) GetRestaurants() (*mongo.Cursor, error) {

	restaurantCollection := db.Client.Database("master_new_copied").Collection("restaurants")

	cursor, err := restaurantCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	db.Kafka.PublishMessage(&kafka.PublishData{
		EventType: "get_all",
	})

	return cursor, nil
}

func (db *Database) AddRestaurant(in *model.Restaurant) (string, error) {

	// restaurantCollection := db.Client.Database("master_new_copied").Collection("restaurants")

	var inerstionID string
	database := db.Database("master_new_copied")
	client := database.Client()
	restaurantCollection := database.Collection("restaurants")

	// start the session
	session, err := client.StartSession()
	if err != nil {
		return "", err
	}
	defer session.EndSession(context.Background())

	_, err = session.WithTransaction(context.Background(), func(sessionContext mongo.SessionContext) (interface{}, error) {
		result, err := restaurantCollection.InsertOne(sessionContext, in)
		if err != nil {
			return "", err
		}

		if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
			inerstionID = oid.Hex()
		}

		return "", nil
	})

	return inerstionID, err
}

func (db *Database) GetRestaurant(id *primitive.ObjectID) (*model.Restaurant, error) {
	restaurantCollection := db.Client.Database("master_new_copied").Collection("restaurants")

	var restaurant model.Restaurant
	result := restaurantCollection.FindOne(context.Background(), &bson.M{"_id": id})
	if err := result.Decode(&restaurant); err != nil {
		return nil, err
	}

	return &restaurant, nil
}

func (db *Database) UpdateRestaurant(id *primitive.ObjectID, in *bson.M) (string, error) {
	restaurantCollection := db.Client.Database("master_new_copied").Collection("restaurants")

	result := restaurantCollection.FindOneAndUpdate(context.Background(), bson.M{"_id": id}, bson.M{"$set": in}, options.FindOneAndUpdate().SetReturnDocument(1))

	if result.Err() != nil {
		return "", result.Err()
	}

	var doc model.Restaurant
	if err := result.Decode(&doc); err != nil {
		return "", err
	}

	return doc.ID.Hex(), nil
}
