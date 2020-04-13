package api

import (
	"context"
	"fmt"

	restaurant "github.com/rohan-luthra/protorepo-restaurants-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *API) AddTax(ctx context.Context, in *restaurant.Tax) (*restaurant.Id, error) {
	return nil, status.Error(codes.Unimplemented, fmt.Sprintf("Working on this :D"))

}
func (a *API) UpdateTax(ctx context.Context, in *restaurant.Tax) (*restaurant.Tax, error) {
	return nil, status.Error(codes.Unimplemented, fmt.Sprintf("Working on this :D"))

}
func (a *API) GetTax(ctx context.Context, in *restaurant.TaxId) (*restaurant.Tax, error) {
	return nil, status.Error(codes.Unimplemented, fmt.Sprintf("Working on this :D"))

}
func (a *API) GetRestTax(restId *restaurant.RestId, stream restaurant.TaxService_GetRestTaxServer) error {
	return status.Error(codes.Unimplemented, fmt.Sprintf("Working on this :D"))

}
func (a *API) DeleteTax(ctx context.Context, in *restaurant.Id) (*restaurant.TaxId, error) {
	return nil, status.Error(codes.Unimplemented, fmt.Sprintf("Working on this :D"))

}
