package app

import (
	"github.com/rohan-luthra/microservice-grpc-go/service-restaurants-go/src/db"
	"github.com/rohan-luthra/microservice-grpc-go/service-restaurants-go/src/kafka"
	"github.com/sirupsen/logrus"
)

// Context App Context
type Context struct {
	Logger        logrus.FieldLogger
	RemoteAddress string
	Database      *db.Database
	Kafka         *kafka.Kafka
}

// WithLogger logger context
func (ctx *Context) WithLogger(logger logrus.FieldLogger) *Context {
	ret := *ctx
	ret.Logger = logger
	return &ret
}

// WithRemoteAddress remote address
func (ctx *Context) WithRemoteAddress(address string) *Context {
	ret := *ctx
	ret.RemoteAddress = address
	return &ret
}
