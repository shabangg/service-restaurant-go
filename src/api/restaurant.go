package api

import (
	"context"
	"fmt"

	"github.com/rohan-luthra/microservice-grpc-go/service-restaurants-go/src/model"
	restaurant "github.com/rohan-luthra/protorepo-restaurants-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *API) GetRestaurants(req *restaurant.Empty, stream restaurant.RestaurantService_GetRestaurantsServer) error {

	ctx := a.App.NewContext()
	result := &model.Restaurant{}

	cursor, err := ctx.GetRestaurants()
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {

		// Decode the data at the current pointer and write it to data
		err := cursor.Decode(result)
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}

		restaurantResponse := result.ToRestaurant()

		stream.Send(restaurantResponse)
	}

	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
	}

	return nil
}

func (a *API) AddRestaurant(ctx context.Context, req *restaurant.Restaurant) (*restaurant.Id, error) {

	newRestaurant := model.Restaurant{
		Name: model.Name{
			En: req.Name.GetEn(),
			Ja: req.Name.GetJa(),
		},
		Username:        req.Username,
		Hashword:        req.Hashword,
		HashwordSalt:    req.HashwordSalt,
		PersonOfContact: req.PersonOfContact,
		Logo:            req.Logo,
		Active:          req.Active,
		Images:          req.Images,
	}

	context := a.App.NewContext()
	result, err := context.AddRestaurant(&newRestaurant)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Invalid Document: %v", err),
		)
	}

	return &restaurant.Id{Id: result}, nil

}

func (a *API) GetRestaurant(ctx context.Context, in *restaurant.Id) (*restaurant.Restaurant, error) {
	context := a.App.NewContext()

	id, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}

	result, err := context.GetRestaurant(&id)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find blog with Object Id %s: %v", in.GetId(), err))
	}

	restaurantResponse := result.ToRestaurant()

	return restaurantResponse, nil
}

func (a *API) UpdateRestaurant(ctx context.Context, in *restaurant.Restaurant) (*restaurant.Id, error) {
	id, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert the supplied blog id to a MongoDB ObjectId: %v", err))
	}

	restaurantData := bson.M{
		"name": in.GetName(),
	}

	context := a.App.NewContext()
	result, err := context.UpdateRestaurant(&id, &restaurantData)

	if err != nil {
		return nil, err
	}

	return &restaurant.Id{Id: result}, nil
}

func (a *API) DeleteRestaurant(ctx context.Context, id *restaurant.Id) (*restaurant.Id, error) {

	// context := a.App.NewContext()

	return nil, status.Error(codes.Unimplemented, fmt.Sprintf("Working on this :D"))
}
