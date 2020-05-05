package db

import (
	"context"
	"fmt"
	"log"

	"github.com/rohan-luthra/microservice-grpc-go/service-restaurants-go/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db *Database) GetRestaurants() (*mongo.Cursor, error) {

	cursor, err := db.RestaurantCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	// db.Kafka.PublishMessage(&kafka.PublishData{
	// 	EventType: "get_all",
	// })

	return cursor, nil
}

func (db *Database) AddRestaurant(in *model.Restaurant) (string, error) {

	var InsertedID string
	result, err := db.RestaurantCollection.InsertOne(context.Background(), in)
	if err != nil {
		return "", err
	}
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		InsertedID = oid.Hex()
	}

	return InsertedID, err
}

func (db *Database) GetRestaurant(id *primitive.ObjectID) (*model.Restaurant, error) {
	var restaurant model.Restaurant

	fmt.Println(id)

	result := db.RestaurantCollection.FindOne(context.Background(), &bson.M{"_id": id})
	fmt.Println(result)
	if err := result.Decode(&restaurant); err != nil {
		log.Panicf("mongo db error %v", err)
		return nil, err
	}

	return &restaurant, nil
}

func (db *Database) UpdateRestaurant(id *primitive.ObjectID, in *bson.M) (string, error) {
	result := db.RestaurantCollection.FindOneAndUpdate(context.Background(), bson.M{"_id": id}, bson.M{"$set": in}, options.FindOneAndUpdate().SetReturnDocument(1))

	if result.Err() != nil {
		return "", result.Err()
	}

	var doc model.Restaurant
	if err := result.Decode(&doc); err != nil {
		return "", err
	}

	return doc.ID.Hex(), nil
}
