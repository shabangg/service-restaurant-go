package app

import (
	"github.com/rohan-luthra/microservice-grpc-go/service-restaurants-go/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ReadCompanies creates a company
func (ctx *Context) GetRestaurants() (*mongo.Cursor, error) {
	return ctx.Database.GetRestaurants()
}

func (ctx *Context) AddRestaurant(restauarnt *model.Restaurant) (string, error) {
	return ctx.Database.AddRestaurant(restauarnt)
}

func (ctx *Context) GetRestaurant(id *primitive.ObjectID) (*model.Restaurant, error) {
	return ctx.Database.GetRestaurant(id)
}

func (ctx *Context) UpdateRestaurant(id *primitive.ObjectID, restauarnt *bson.M) (string, error) {
	return ctx.Database.UpdateRestaurant(id, restauarnt)
}
