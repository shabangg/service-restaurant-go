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

var OrderType_name = restaurant.OrderType_name

type Tax struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	RestaurantId primitive.ObjectID `bson:"rest_id,omitempty"`
	Name         Name               `bson:"name,omitempty"`
	Inclusive    bool               `bson:"inclusive,omitempty"`
	IsPercentage bool               `bson:"is_percentage,omitempty"`
	Value        float32            `bson:"value,omitempty"`
	OrderTypes   []string           `bson:"order_types,omitempty"`
	CreatedAt    time.Time          `bson:"created_at,omitempty"`
	UpdatedAt    time.Time          `bson:"Updated_at,omitempty"`
}

func (model *Tax) ToTax() (*restaurant.Tax, error) {

	createdAt, err := ptypes.TimestampProto(model.CreatedAt)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
	}
	updatedAt, err := ptypes.TimestampProto(model.UpdatedAt)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
	}

	return &restaurant.Tax{
		Id: model.Id.Hex(),
		Name: &restaurant.Name{
			En: model.Name.En,
			Ja: model.Name.Ja,
		},
		RestId:       model.RestaurantId.Hex(),
		Inclusive:    model.Inclusive,
		IsPercentage: model.IsPercentage,
		Value:        model.Value,
		// OrderTypes:   OrderType_name.,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil

}
