package catalog_service

import (
	"context"

	"github.com/minhhoanq/lifeat/catalog_service/internal/handler/grpc"
	"github.com/minhhoanq/lifeat/common/logger"
)

type Server struct {
	grpcServer grpc.Server
	l          logger.Interface
}

func NewServer(grpcServer grpc.Server, l logger.Interface) *Server {
	return &Server{
		grpcServer: grpcServer,
		l:          l,
	}
}

func (s *Server) Start() {
	s.grpcServer.Start(context.Background())
	s.l.Info("Start gRPC server successfully")
}
