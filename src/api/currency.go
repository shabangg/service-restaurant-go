package api

import (
	"context"
	"fmt"

	restaurant "github.com/rohan-luthra/protorepo-restaurants-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *API) AddCurrency(ctx context.Context, in *restaurant.Currency) (*restaurant.Id, error) {
	return nil, status.Error(codes.Unimplemented, fmt.Sprintf("Working on this :D"))

}
func (a *API) UpdateCurrency(ctx context.Context, in *restaurant.Currency) (*restaurant.Currency, error) {
	return nil, status.Error(codes.Unimplemented, fmt.Sprintf("Working on this :D"))

}
func (a *API) GetCurrencies(restId *restaurant.RestId, stream restaurant.CurrencyService_GetCurrenciesServer) error {
	return status.Error(codes.Unimplemented, fmt.Sprintf("Working on this :D"))

}
func (a *API) DeleteCurrency(ctx context.Context, in *restaurant.Id) (*restaurant.CurrencyId, error) {
	return nil, status.Error(codes.Unimplemented, fmt.Sprintf("Working on this :D"))

}
