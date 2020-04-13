package model

import (
	restaurant "github.com/rohan-luthra/protorepo-restaurants-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Name struct {
	En string `bson:"en,omitempty"`
	Ja string `bson:"ja,omitempty"`
}

type Contact struct {
	PhoneNumber string `bson:"phone_number,omitempty"`
	Email       string `bson:"email,omitempty"`
	PersonName  string `bson:"person_name,omitempty"`
}

type Slots struct {
	StartTime string `bson:"start_time,omitempty"`
	EndTime   string `bson:"end_time,omitempty"`
}

type Timings struct {
	Day   string  `bson:"start_time,omitempty"`
	Slots []Slots `bson:"slots,omitempty"`
}

type Restaurant struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Name              Name               `bson:"name,omitempty"`
	Username          string             `bson:"username,omitempty"`
	Hashword          string             `bson:"hashword,omitempty"`
	HashwordSalt      int32              `bson:"hashword_salt,omitempty"`
	PersonOfContact   string             `bson:"person_of_contact,omitempty"`
	Logo              string             `bson:"logo,omitempty"`
	Active            bool               `bson:"active,omitempty"`
	Images            []string           `bson:"images,omitempty"`
	Contacts          []Contact          `bson:"contacts,omitempty"`
	PaymentModes      []string           `bson:"payment_modes,omitempty"`
	Timings           []Timings          `bson:"timings,omitempty"`
	SubscriptionPlan  string             `bson:"subscription_plan,omitempty"`
	SubscriptionPrice float32            `bson:"subscription_price,omitempty"`
	// TODO: all Currency
	TrailDays float32 `bson:"trial_period,omitempty"`
}

func (model *Restaurant) ToRestaurant() *restaurant.Restaurant {

	return &restaurant.Restaurant{
		Id: model.ID.Hex(),
		Name: &restaurant.Name{
			En: model.Name.En,
			Ja: model.Name.Ja,
		},
		Username:     model.Username,
		Hashword:     model.Hashword,
		HashwordSalt: model.HashwordSalt,
		// ContactNumber: model.ContactNumber,

	}
}
