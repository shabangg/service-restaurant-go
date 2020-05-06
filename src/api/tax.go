package api

import (
	"context"
	"fmt"

	restaurant "github.com/rohan-luthra/protorepo-restaurants-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AddTax - add a new Tax
func (a *API) AddTax(ctx context.Context, in *restaurant.Tax) (*restaurant.Id, error) {
	return nil, status.Error(codes.Unimplemented, fmt.Sprintf("Working on this :D"))

}

// UpdateTax - update en existing Tax
func (a *API) UpdateTax(ctx context.Context, in *restaurant.Tax) (*restaurant.Tax, error) {
	return nil, status.Error(codes.Unimplemented, fmt.Sprintf("Working on this :D"))

}

// GetTax -get a Tax of id
func (a *API) GetTax(ctx context.Context, in *restaurant.TaxId) (*restaurant.Tax, error) {
	return nil, status.Error(codes.Unimplemented, fmt.Sprintf("Working on this :D"))

}

// GetRestTax -  get all Tax for a restaurant_id
func (a *API) GetRestTax(restId *restaurant.RestId, stream restaurant.TaxService_GetRestTaxServer) error {
	return status.Error(codes.Unimplemented, fmt.Sprintf("Working on this :D"))

}

// DeleteTax - delete a Tax of id
func (a *API) DeleteTax(ctx context.Context, in *restaurant.Id) (*restaurant.TaxId, error) {
	return nil, status.Error(codes.Unimplemented, fmt.Sprintf("Working on this :D"))

}
