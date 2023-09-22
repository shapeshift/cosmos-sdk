package grpc

import (
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"

	"github.com/shapeshift/cosmos-sdk/client"
	"github.com/shapeshift/cosmos-sdk/codec"
	"github.com/shapeshift/cosmos-sdk/server/config"
	"github.com/shapeshift/cosmos-sdk/server/grpc/gogoreflection"
	reflection "github.com/shapeshift/cosmos-sdk/server/grpc/reflection/v2alpha1"
	"github.com/shapeshift/cosmos-sdk/server/types"
	sdk "github.com/shapeshift/cosmos-sdk/types"
	_ "github.com/shapeshift/cosmos-sdk/types/tx/amino" // Import amino.proto file for reflection
)

// StartGRPCServer starts a gRPC server on the given address.
func StartGRPCServer(clientCtx client.Context, app types.Application, cfg config.GRPCConfig) (*grpc.Server, error) {
	maxSendMsgSize := cfg.MaxSendMsgSize
	if maxSendMsgSize == 0 {
		maxSendMsgSize = config.DefaultGRPCMaxSendMsgSize
	}

	maxRecvMsgSize := cfg.MaxRecvMsgSize
	if maxRecvMsgSize == 0 {
		maxRecvMsgSize = config.DefaultGRPCMaxRecvMsgSize
	}

	grpcSrv := grpc.NewServer(
		grpc.ForceServerCodec(codec.NewProtoCodec(clientCtx.InterfaceRegistry).GRPCCodec()),
		grpc.MaxSendMsgSize(maxSendMsgSize),
		grpc.MaxRecvMsgSize(maxRecvMsgSize),
	)

	app.RegisterGRPCServer(grpcSrv)

	// Reflection allows consumers to build dynamic clients that can write to any
	// Cosmos SDK application without relying on application packages at compile
	// time.
	err := reflection.Register(grpcSrv, reflection.Config{
		SigningModes: func() map[string]int32 {
			modes := make(map[string]int32, len(clientCtx.TxConfig.SignModeHandler().Modes()))
			for _, m := range clientCtx.TxConfig.SignModeHandler().Modes() {
				modes[m.String()] = (int32)(m)
			}
			return modes
		}(),
		ChainID:           clientCtx.ChainID,
		SdkConfig:         sdk.GetConfig(),
		InterfaceRegistry: clientCtx.InterfaceRegistry,
	})
	if err != nil {
		return nil, err
	}

	// Reflection allows external clients to see what services and methods
	// the gRPC server exposes.
	gogoreflection.Register(grpcSrv)

	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		return nil, err
	}

	errCh := make(chan error)
	go func() {
		err = grpcSrv.Serve(listener)
		if err != nil {
			errCh <- fmt.Errorf("failed to serve: %w", err)
		}
	}()

	select {
	case err := <-errCh:
		return nil, err

	case <-time.After(types.ServerStartTime):
		// assume server started successfully
		return grpcSrv, nil
	}
}
