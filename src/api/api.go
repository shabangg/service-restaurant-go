package api

import "github.com/rohan-luthra/microservice-grpc-go/service-restaurants-go/src/app"

// API api base struct
type API struct {
	App    *app.App
	Config *Config
}

// New new api instance
func New(a *app.App) (api *API, err error) {
	api = &API{App: a}

	api.Config, err = InitConfig()
	if err != nil {
		return nil, err
	}

	return api, nil
}
