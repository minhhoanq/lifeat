package grpc

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/lib/pq"
	pbuser "github.com/minhhoanq/lifeat/user_service/internal/controller/grpc/v1/user_service"
	"github.com/minhhoanq/lifeat/user_service/internal/entity"
	"github.com/minhhoanq/lifeat/user_service/internal/usecase/rest/repo"
	"github.com/minhhoanq/lifeat/user_service/internal/worker"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *GrpcServer) Signup(ctx context.Context, arg *pbuser.SignupRequest) (*pbuser.SignupResponse, error) {
	hashedPassword, err := entity.HashPassword(arg.Password)
	if err != nil {
		return nil, err
	}

	args := repo.CreateUserTxParams{
		CreateUserParams: repo.CreateUserParams{
			Username: arg.Username,
			Email:    arg.Email,
			Password: hashedPassword,
			RoleId:   1,
		},
		AfterCreate: func(user *entity.User) error {
			taskPayload := &worker.PayloadSendVerifyEmail{UserId: user.ID}

			opts := []asynq.Option{
				asynq.MaxRetry(10),
				asynq.Queue(worker.QueueCritial),
			}

			return server.taskDistributor.DistributeTaskSendVerifyEmail(ctx, taskPayload, opts...)
		},
	}

	txResult, err := server.q.CreateUserTx(ctx, args)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "username already exists %s", err)

			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create user %s", err.Error())
	}

	rsp := &pbuser.SignupResponse{
		User: convertUser(*txResult.User),
	}

	return rsp, nil
}
