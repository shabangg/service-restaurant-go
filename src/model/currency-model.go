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

type Currency struct {
	Id     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string             `bson:"name,omitempty"`
	Symbol string             `bson:"symbol,omitempty"`

	CreatedAt time.Time `bson:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"Updated_at,omitempty"`
}

func (model *Currency) ToCurrency() (*restaurant.Currency, error) {

	createdAt, err := ptypes.TimestampProto(model.CreatedAt)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
	}
	updatedAt, err := ptypes.TimestampProto(model.UpdatedAt)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
	}

	return &restaurant.Currency{
		Id:        model.Id.Hex(),
		Name:      model.Name,
		Symbol:    model.Symbol,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}
