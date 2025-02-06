package worker

import (
	"context"

	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

type TaskProcessor interface {
	ProcessTaskVerifyEmail(context.Context, *asynq.Task) error
	Start() error
	Shutdown()
}

type RedisTaskProcessor struct {
	server *asynq.Server
	db     *gorm.DB
}

// func NewRedisTaskProcessor(redisOpts asynq.RedisClientOpt, db *gorm.DB) TaskProcessor {
// 	server :=
// }
