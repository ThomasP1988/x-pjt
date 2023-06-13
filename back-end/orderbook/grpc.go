package main

import (
	grpc_ob "NFTM/shared/orderbook/grpc"
	"log"
	"net"

	// "orderbook_service/handler/server"
	"NFTM/shared/orderbook/server"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	port string = ":50052"
)

func GenerateTLSApi(pemPath, keyPath string) (*grpc.Server, error) {
	cred, err := credentials.NewServerTLSFromFile(pemPath, keyPath)
	if err != nil {
		return nil, err
	}

	s := grpc.NewServer(
		grpc.Creds(cred),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_zap.StreamServerInterceptor(&zap.Logger{}),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(&zap.Logger{}),
		)),
	)
	return s, nil
}

func StartGRPC() {

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	apiserver := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(grpc_zap.StreamServerInterceptor(logger))),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(logger),
		)),
	)

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpc_ob.RegisterOrderbookServer(apiserver, &server.OrderbookServer{})
	StartStandAlone(apiserver)

	log.Printf("server listening at %v", lis.Addr())
	if err := apiserver.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
