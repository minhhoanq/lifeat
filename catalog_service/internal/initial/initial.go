package initial

import (
	"context"

	"github.com/minhhoanq/lifeat/catalog_service/configs"
	"github.com/minhhoanq/lifeat/catalog_service/internal/dataaccess/database"
	"github.com/minhhoanq/lifeat/catalog_service/internal/dataaccess/redis"
	"github.com/minhhoanq/lifeat/catalog_service/internal/handler/grpc"
	"github.com/minhhoanq/lifeat/catalog_service/internal/service"
	"github.com/minhhoanq/lifeat/common/logger"
)

func InitialServer(cfg configs.Config, l logger.Interface) (grpc.Server, error) {

	db, err := database.New(cfg, l)
	if err != nil {
		return nil, err
	}

	// connect to redis
	redis := redis.NewRedis(cfg, l)
	redisClient := redis.Connect()
	_, err = redisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	catalogAccessor := database.NewCatalogDataAccessor(db, l)
	catalogService := service.NewCatalogService(db.DB, l, catalogAccessor)
	handler, err := grpc.NewHandler(catalogService, l, redisClient)
	if err != nil {
		return nil, err
	}
	server := grpc.NewGRPCServer(cfg, handler, l)

	return server, nil
}
