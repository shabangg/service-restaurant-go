package validate

import (
	"fmt"

	restaurant "github.com/rohan-luthra/protorepo-restaurants-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidateAddRestaurant(r *restaurant.AddRestaurantReq) error {

	if r.Username == "" {
		return status.Errorf(codes.Internal, fmt.Sprint("Missing Username"))
	}
	if r.Password == "" {
		return status.Errorf(codes.Internal, fmt.Sprint("Missing Password"))
	}

	if len(r.Password) >= 6 {
		return status.Errorf(codes.Internal, fmt.Sprint("Password should be atleast 6 characters long"))
	}

	if r.Name == nil {
		return status.Errorf(codes.Internal, fmt.Sprint("Missing Name"))
	}

	if r.Name.En == "" {
		return status.Errorf(codes.Internal, fmt.Sprint("Missing Name field En"))
	}

	if r.DefaultCurrenyId == "" {
		return status.Errorf(codes.Internal, fmt.Sprint("Missing Default Currency Id"))
	}

	_, err := primitive.ObjectIDFromHex(r.DefaultCurrenyId)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, fmt.Sprint("Invalid default_currency_id in Address"))
	}

	if err = ValidateAddress(r.Address); err != nil {
		return err
	}

	return nil
}

func ValidateAddress(a *restaurant.Address) error {

	if a == nil {
		return status.Errorf(codes.Internal, fmt.Sprint("Missing Address"))
	}
	if a.Line1 == "" || a.CityId == "" || a.CountryId == "" || a.Pincode == 0 {
		return status.Errorf(codes.Internal, fmt.Sprint("Missing Address Fields"))
	}

	_, err := primitive.ObjectIDFromHex(a.CityId)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, fmt.Sprint("Invalid city_id in Address"))
	}

	_, err = primitive.ObjectIDFromHex(a.CountryId)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, fmt.Sprint("Invalid country_id in Address"))
	}

	return nil
}
