package redis

import (
	"context"
	"time"

	"github.com/minhhoanq/lifeat/common/logger"
	"github.com/minhhoanq/lifeat/order_service/config"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Redis struct {
	cfg config.Config
	l   logger.Interface
}

func NewRedis(cfg config.Config, l logger.Interface) *Redis {
	return &Redis{
		cfg: cfg,
		l:   l,
	}
}

func (r *Redis) Connect() *redis.Client {
	// Connect to Redis
	opts := &redis.Options{
		Addr: r.cfg.RedisAddress,
	}
	client := redis.NewClient(opts)

	r.l.Info("connecting to redis", logger.String("address", r.cfg.RedisAddress))

	return client
}

func (r *Redis) Close() {
	r.Close()
}

func (r *Redis) AcquireLock(client *redis.Client, lockKey string, timeout time.Duration) bool {
	ctx := context.Background()

	// Try to acquire the lock with SETNX command (SET if Not Exists)
	lockAcquire, err := client.SetNX(ctx, lockKey, "1", timeout).Result()
	if err != nil {
		r.l.Error("error acquiring lock: ", zap.Error(err))
		return false
	}

	return lockAcquire
}

func ReleaseLock(client *redis.Client, lockKey string) {
	ctx := context.Background()
	client.Del(ctx, lockKey)
}
