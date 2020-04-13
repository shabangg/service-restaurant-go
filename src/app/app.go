package app

import (
	"context"

	"github.com/rohan-luthra/microservice-grpc-go/service-restaurants-go/src/db"
	"github.com/rohan-luthra/microservice-grpc-go/service-restaurants-go/src/kafka"
	"github.com/sirupsen/logrus"
)

// App Application
type App struct {
	Config   *Config
	Database *db.Database
	Kafka    *kafka.Kafka
}

// NewContext return new application context
func (a *App) NewContext() *Context {
	return &Context{
		Logger:   logrus.StandardLogger(),
		Database: a.Database,
		Kafka:    a.Kafka,
	}
}

// New Application new instance
func New() (app *App, err error) {
	app = &App{}

	app.Config, err = InitConfig()
	if err != nil {
		return nil, err
	}

	err, dbConfig := db.InitConfig()
	if err != nil {
		return nil, err
	}

	err, app.Database = db.New(dbConfig)
	if err != nil {
		return nil, err
	}

	return app, err
}

// Close close the database
func (app *App) Close() error {
	return app.Database.Disconnect(context.TODO())
}
