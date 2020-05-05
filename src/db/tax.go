package db

import (
	"context"

	"github.com/rohan-luthra/microservice-grpc-go/service-restaurants-go/src/model"
	restaurant "github.com/rohan-luthra/protorepo-restaurants-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db *Database) GetTax(id *primitive.ObjectID) (*model.Tax, error) {
	var tax model.Tax
	result := db.TaxCollection.FindOne(context.Background(), &bson.M{"_id": id})
	if err := result.Decode(&tax); err != nil {
		return nil, err
	}
	return &tax, nil
}

func (db *Database) GetRestTax(restId *restaurant.RestId) (*mongo.Cursor, error) {

	cursor, err := db.TaxCollection.Find(context.Background(), bson.M{"rest_id": restId})
	if err != nil {
		return nil, err
	}
	return cursor, nil
}

func (db *Database) AddTax(in *model.Tax) (string, error) {

	var inerstionID string
	// start the session
	session, err := db.Client.StartSession()
	if err != nil {
		return "", err
	}
	defer session.EndSession(context.Background())

	_, err = session.WithTransaction(context.Background(), func(sessionContext mongo.SessionContext) (interface{}, error) {
		result, err := db.TaxCollection.InsertOne(sessionContext, in)
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

func (db *Database) UpdateTax(id *primitive.ObjectID, in *bson.M) (*model.Tax, error) {

	result := db.RestaurantCollection.FindOneAndUpdate(context.Background(), bson.M{"_id": id}, bson.M{"$set": in}, options.FindOneAndUpdate().SetReturnDocument(1))

	if result.Err() != nil {
		return nil, result.Err()
	}

	var doc model.Tax
	if err := result.Decode(&doc); err != nil {
		return nil, err
	}

	return &doc, nil
}
