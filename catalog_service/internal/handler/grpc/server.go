package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/minhhoanq/lifeat/order_service/internal/configs"
	pb "github.com/minhhoanq/lifeat/order_service/internal/generated/catalog_service"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/validator"
	"google.golang.org/grpc"
)

type Server interface {
	Start(ctx context.Context) error
}

type server struct {
	grpcConfig configs.GRPC
	handler    pb.CatalogServiceServer
}

func NewServer(
	grpcConfig configs.GRPC,
	handler pb.CatalogServiceServer,
) Server {
	return &server{
		grpcConfig: grpcConfig,
		handler:    handler,
	}
}

func (s *server) Start(ctx context.Context) error {
	listener, err := net.Listen("tcp", s.grpcConfig.Address)
	if err != nil {
		return err
	}
	defer listener.Close()

	var opts = []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			validator.UnaryServerInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			validator.StreamServerInterceptor(),
		),
	}
	server := grpc.NewServer(opts...)
	pb.RegisterCatalogServiceServer(server, s.handler)
	fmt.Printf("gRPC server is running on %s\n", s.grpcConfig.Address)
	return server.Serve(listener)
}
