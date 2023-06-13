package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	grpc_ob "NFTM/apis/grpc/expose/orderbook"
	ob_handler "NFTM/apis/grpc/handler/orderbook"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
)

const (
	portGRPC    = ":50051"
	portGRPCWeb = ":50052"
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

func GetGRPC() *grpc.Server {
	apiserver := grpc.NewServer()
	grpc_ob.RegisterOrderbooksServer(apiserver, ob_handler.NewOrderbookServer())
	return apiserver
}

func StartGRPC() {
	apiserver := GetGRPC()
	lis, err := net.Listen("tcp", portGRPC)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("GRPC server listening at %v", lis.Addr())

	if err := apiserver.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func GetWrappedServer(grpcServer *grpc.Server) *grpcweb.WrappedGrpcServer {
	return grpcweb.WrapServer(grpcServer, grpcweb.WithOriginFunc(func(origin string) bool { return true }), grpcweb.WithAllowedRequestHeaders([]string{"*"}))
}

func StartGRPCHTTP() {
	grpcServer := GetGRPC()

	wrappedGrpc := GetWrappedServer(grpcServer)

	tlsHttpServer := http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// fmt.Printf("req.RequestURI: %v\n", r.RequestURI)
			fmt.Printf("req.RequestURI: %v\n", r.RequestURI)
			// fmt.Printf("r.Body: %v\n", r.Body)
			// fmt.Printf("r.Method: %v\n", r.Method)

			if wrappedGrpc.IsGrpcWebSocketRequest(r) {
				fmt.Printf("\"websocket\": %v\n", "websocket")

				wrappedGrpc.HandleGrpcWebsocketRequest(w, r)
			}

			if wrappedGrpc.IsGrpcWebRequest(r) || wrappedGrpc.IsAcceptableGrpcCorsRequest(r) {

				// fmt.Printf("GRPC-WEB %+v", r)
				wrappedGrpc.ServeHTTP(w, r)

				return
			}
			// Fall back to other servers.
			http.DefaultServeMux.ServeHTTP(w, r)
		}),
	}

	lis, err := net.Listen("tcp", portGRPCWeb)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("GRPC Web HTTP server listening at %v", lis.Addr())
	tlsHttpServer.Serve(lis)
}
