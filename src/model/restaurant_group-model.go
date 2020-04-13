package model

import (
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"
	restaurant "github.com/rohan-luthra/protorepo-restaurants-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RestaurantGroup struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	Username     string             `bson:"username,omitempty"`
	Hashword     string             `bson:"hashword,omitempty"`
	HashwordSalt int32              `bson:"hasword-salt,omitempty"`
	Contact      []Contact          `bson:"contacts,omitempty"`
	Restaurants  []string           `bson:"restaurants,omitempty`

	CreatedAt time.Time `bson:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"Updated_at,omitempty"`
}

func (model *RestaurantGroup) ToRestaurantGroup() (*restaurant.RestaurantGroup, error) {

	createdAt, err := ptypes.TimestampProto(model.CreatedAt)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
	}
	updatedAt, err := ptypes.TimestampProto(model.UpdatedAt)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
	}

	return &restaurant.RestaurantGroup{
		Id:           model.Id.Hex(),
		Username:     model.Username,
		Hasword:      model.Hashword,
		HashwordSalt: model.HashwordSalt,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}, nil
}
