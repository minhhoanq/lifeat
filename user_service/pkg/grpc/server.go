package grpc

import (
	"context"

	"github.com/minhhoanq/lifeat/user_service/config"
	pbuser "github.com/minhhoanq/lifeat/user_service/internal/controller/grpc/v1/api/v1/user_service"
)

const ()

type GrpcServer struct {
	pbuser.UnimplementedUserServiceServer
	cfg config.Config
	ctx context.Context
}

func NewRestServer(cfg config.Config, ctx context.Context) (*GrpcServer, error) {
	return &GrpcServer{
		cfg: cfg,
		ctx: ctx,
	}, nil
}
