package test

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"testing"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	ratelimit "github.com/grpc-ecosystem/go-grpc-middleware/ratelimit"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/rohan-luthra/microservice-grpc-go/service-restaurants-go/src/api"
	"github.com/rohan-luthra/microservice-grpc-go/service-restaurants-go/src/app"
	"github.com/spf13/viper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"

	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip" // Install the gzip compressor
	"google.golang.org/grpc/keepalive"

	restaurant "github.com/rohan-luthra/protorepo-restaurants-go"
)

var configFile string

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {

	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.AddConfigPath("./..")
		viper.SetConfigName("config_test")
		viper.SetConfigType("yml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("unable to read config: %v\n", err)
		os.Exit(1)
	}
}

func TestImplementation_GRPCRouting_GetPositionByID_ShouldOK(t *testing.T) {

	ctx := context.Background()
	srv, listener := startGRPCServer()

	// it is here to properly stop the server
	defer func() { time.Sleep(10 * time.Millisecond) }()
	defer srv.Stop()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	// defer conn.Close()
	// log.Println(conn)

	c := restaurant.NewRestaurantServiceClient(conn)
	log.Printf("%v", c)
	rest, err := c.GetRestaurant(ctx, &restaurant.Id{Id: "5e561492d08bbb00174fc4ec"})
	if err != nil {
		t.Fatalf("error %v", err)
	}
	log.Println(rest)
}

type alwaysPassLimiter struct{}

func (*alwaysPassLimiter) Limit() bool {
	return false
}

func startGRPCServer() (*grpc.Server, *bufconn.Listener) {
	bufferSize := 1024 * 1024
	listener := bufconn.Listen(bufferSize)

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
	a, err := app.New()
	if err != nil {
		log.Fatalf("app error %v", err)
	}
	defer a.Close()

	api, err := api.New(a)
	if err != nil {
		log.Fatalf("api error %v", err)
	}
	restaurant.RegisterRestaurantServiceServer(grpcServer, api)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to start grpc server: %v", err)
		}
	}()
	return grpcServer, listener
}

func getBufDialer(listener *bufconn.Listener) func(context.Context, string) (net.Conn, error) {
	return func(ctx context.Context, url string) (net.Conn, error) {
		return listener.Dial()
	}
}
