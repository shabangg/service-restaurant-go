package model

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	restaurant "github.com/rohan-luthra/protorepo-restaurants-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	Line1        string  `bson:"line1,omitempty"`
	Line2        string  `bson:"line2,omitempty"`
	City         string  `bson:"city,omitempty"`
	CityId       string  `bson:"city_id,omitempty"`
	State        string  `bson:"state,omitempty"`
	Country      string  `bson:"country,omitempty"`
	CountryId    string  `bson:"country_id,omitempty"`
	Pincode      int32   `bson:"pincode,omitempty"`
	GeoLatitude  float32 `bson:"geo_latitude,omitempty"`
	GeoLongitude float32 `bson:"geo_longitude,omitempty"`
	Timezone     string  `bson:"timezone,omitempty"`
}

type Name struct {
	En string `bson:"en,omitempty" validate:"required"`
	Ja string `bson:"ja,omitempty"`
}

type Contact struct {
	PhoneNumber string `bson:"phone_number,omitempty"`
	Email       string `bson:"email,omitempty"`
	Name        string `bson:"name,omitempty"`
}
type Slot struct {
	StartTime string `bson:"start_time,omitempty"`
	EndTime   string `bson:"end_time,omitempty"`
}

type Timing struct {
	Day   string `bson:"start_time,omitempty"`
	Slots []Slot `bson:"slots,omitempty"`
}

type Restaurant struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty" validate:"required"`
	Name               *Name              `bson:"name,omitempty" validate:"required"`
	Username           string             `bson:"username,omitempty"`
	Hashword           string             `bson:"hashword,omitempty"`
	HashwordSalt       int32              `bson:"hashword_salt,omitempty"`
	Contact            *Contact           `bson:"contacts,omitempty"`
	PersonOfContact    string             `bson:"person_of_contact,omitempty"`
	Logo               string             `bson:"logo,omitempty"`
	ProfileImage       string             `bson:"profile_image,omitempty"`
	Active             bool               `bson:"active,omitempty"`
	Address            *Address           `bson:"address,omitempty"`
	Images             []string           `bson:"images,omitempty"`
	FcmTokens          []string           `bson:"fcm_token,omitempty"`
	PaymentMode        []string           `bson:"payment_mode,omitempty"`
	Timings            []*Timing          `bson:"timings,omitempty"`
	SubscriptionPlan   string             `bson:"subscription_plan,omitempty"`
	SubscriptionPrice  float32            `bson:"subscription_price,omitempty"`
	CurrencyId         string
	Curreny            *Currency            `bson:"curreny,omitempty"`
	AssignedSalesEmpId string               `bson:"assigned_sales_emp_id,omitempty"`
	AssignedOpsEmpId   string               `bson:"assigned_ops_emp_id,omitempty"`
	TrialDays          float32              `bson:"trial_days,omitempty"`
	CreatedAt          *timestamp.Timestamp `bson:"created_at,omitempty"`
	UpdatedAt          *timestamp.Timestamp `bson:"updated_at,omitempty"`
}

func ToRestaurant(model *Restaurant) (*restaurant.Restaurant, error) {

	return &restaurant.Restaurant{
		Id:                 model.ID.Hex(),
		Username:           model.Username,
		Hashword:           model.Hashword,
		HashwordSalt:       model.HashwordSalt,
		ProfileImage:       model.ProfileImage,
		Active:             model.Active,
		Images:             model.Images,
		SubscriptionPrice:  model.SubscriptionPrice,
		TrialDays:          model.TrialDays,
		PersonOfContact:    model.PersonOfContact,
		AssignedOpsEmpId:   model.AssignedOpsEmpId,
		AssignedSalesEmpId: model.AssignedSalesEmpId,
		CreatedAt:          model.CreatedAt,
		UpdatedAt:          model.UpdatedAt,
		Logo:               model.Logo,
		FcmTokens:          model.FcmTokens,
		// Timings:            getTimings(model.Timings),
		// PaymentMode:        model.PaymentMode,
		// Curreny:            model.Curreny,
		Address: &restaurant.Address{
			Line1:   model.Address.Line1,
			Line2:   model.Address.Line2,
			City:    model.Address.City,
			Country: model.Address.Country,
			Pincode: model.Address.Pincode,
			State:   model.Address.State,
		},
		Name: &restaurant.Name{
			En: model.Name.En,
			Ja: model.Name.Ja,
		},
	}, nil
}

func getTimings(t []Timing) []*restaurant.Timings {
	var timings []*restaurant.Timings
	for _, elem := range t {
		timings = append(timings, &restaurant.Timings{
			Day:   1,
			Slots: getSlots(elem.Slots),
		})
	}
	return timings
}

// func getDay(t Timing) *restaurant.Days {

// }

func getSlots(s []Slot) []*restaurant.Slot {
	var slots []*restaurant.Slot
	for _, elem := range s {
		slots = append(slots, &restaurant.Slot{
			StartTime: elem.StartTime,
			EndTime:   elem.EndTime,
		})
	}
	return slots
}
