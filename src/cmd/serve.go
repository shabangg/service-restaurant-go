package cmd

import (
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	ratelimit "github.com/grpc-ecosystem/go-grpc-middleware/ratelimit"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/rohan-luthra/microservice-grpc-go/service-restaurants-go/src/api"
	"github.com/rohan-luthra/microservice-grpc-go/service-restaurants-go/src/app"

	restaurant "github.com/rohan-luthra/protorepo-restaurants-go"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip" // Install the gzip compressor
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	"github.com/spf13/cobra"
)

// alwaysPassLimiter is an example limiter which implements Limiter interface.
// It does not limit any request because Limit function always returns false.
type alwaysPassLimiter struct{}

func (*alwaysPassLimiter) Limit() bool {
	return false
}
func listenGRPC(api *api.API, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	recoveryFunc := func(p interface{}) (err error) {
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}

	recoveryOpts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(recoveryFunc),
	}

	// Create unary/stream rateLimiters, based on token bucket here.
	limiter := &alwaysPassLimiter{}

	// keep alive policy
	kaep := keepalive.EnforcementPolicy{
		MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
		PermitWithoutStream: true,            // Allow pings even when there are no active streams
	}
	kasp := keepalive.ServerParameters{
		MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
		MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
		MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
		Time:                  5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
		Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
	}

	grpcServer := grpc.NewServer(
		grpc.KeepaliveEnforcementPolicy(kaep),
		grpc.KeepaliveParams(kasp),
		grpc.ConnectionTimeout(time.Minute*30),
		grpc.MaxRecvMsgSize(1024*1024*128),
		grpc_middleware.WithUnaryServerChain(
			ratelimit.UnaryServerInterceptor(limiter),
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_recovery.UnaryServerInterceptor(recoveryOpts...),
		),
		grpc_middleware.WithStreamServerChain(
			ratelimit.StreamServerInterceptor(limiter),
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_recovery.StreamServerInterceptor(recoveryOpts...),
		),
	)

	restaurant.RegisterRestaurantServiceServer(grpcServer, api)
	restaurant.RegisterTaxServiceServer(grpcServer, api)
	// restaurant.RegisterCurrencyServiceServer(grpcServer, api)
	reflection.Register(grpcServer)

	logrus.Infof("starting HTTP/2 gRPC API server: %q\n", lis.Addr().String())

	return grpcServer.Serve(lis)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serves the gRPC server",
	Long:  `start the gRPC server on provided port`,
	RunE: func(cmd *cobra.Command, args []string) error {

		a, err := app.New()
		if err != nil {
			return err
		}
		defer a.Close()

		api, err := api.New(a)
		if err != nil {
			return err
		}

		port := api.Config.Port

		if err := listenGRPC(api, port); err != nil {
			logrus.Error("Failed to serve: %v\n", err)
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}
