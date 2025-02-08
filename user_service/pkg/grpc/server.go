package grpc

import (
	"context"

	"github.com/minhhoanq/lifeat/user_service/config"
	pbuser "github.com/minhhoanq/lifeat/user_service/internal/controller/grpc/v1/user_service"
	"github.com/minhhoanq/lifeat/user_service/internal/token"
	"github.com/minhhoanq/lifeat/user_service/internal/usecase/rest/repo"
	"github.com/minhhoanq/lifeat/user_service/internal/worker"
)

type GrpcServer struct {
	pbuser.UnimplementedUserServiceServer
	cfg             config.Config
	ctx             context.Context
	taskDistributor worker.TaskDistributor
	q               repo.Querier
	tokenMaker      token.Maker
}

func NewGrpcServer(cfg config.Config, ctx context.Context, taskDistributor worker.TaskDistributor, q repo.Querier, tokenMaker token.Maker) (*GrpcServer, error) {
	return &GrpcServer{
		cfg:             cfg,
		ctx:             ctx,
		taskDistributor: taskDistributor,
		q:               q,
		tokenMaker:      tokenMaker,
	}, nil
}
