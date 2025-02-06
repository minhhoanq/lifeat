package rest

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/hibiken/asynq"
	"github.com/minhhoanq/lifeat/common/logger"
	"github.com/minhhoanq/lifeat/user_service/config"
	v1 "github.com/minhhoanq/lifeat/user_service/internal/controller/rest/v1"
	"github.com/minhhoanq/lifeat/user_service/internal/email"
	"github.com/minhhoanq/lifeat/user_service/internal/token"
	"github.com/minhhoanq/lifeat/user_service/internal/usecase"
	"github.com/minhhoanq/lifeat/user_service/internal/usecase/repo"
	"github.com/minhhoanq/lifeat/user_service/internal/worker"
	"github.com/minhhoanq/lifeat/user_service/pkg/postgres"
	"github.com/minhhoanq/lifeat/user_service/pkg/rest"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

const (
	publicKeyPath  = "config/keys/public.pem"
	privateKeyPath = "config/keys/private.pem"
)

func RunRestServer(cfg config.Config) {
	l := logger.NewWrapLogger(zap.DebugLevel, false)

	pg, err := postgres.New(cfg, l)
	if err != nil {
		l.Error("cannot connection to database", zap.String("Error: ", err.Error()))
		return
	}
	defer pg.Close(l)

	tokenMaker, err := token.NewJWTMaker(publicKeyPath, privateKeyPath)
	if err != nil {
		l.Error("jwt maker error", zap.String("Error: ", err.Error()))
		return
	}

	redisOpts := asynq.RedisClientOpt{
		Addr: cfg.RedisAddres,
	}

	taskDistributor := worker.NewRedisTaskDistributor(l, redisOpts)

	q := repo.New(pg.DB)

	runTaskProcessor(context.Background(), cfg, redisOpts, q, l)
	// Resful
	handler := echo.New()
	u := usecase.New(q, tokenMaker, cfg, taskDistributor)
	v1.NewRouter(handler, l, u, tokenMaker)
	// , rest.Port(cfg.HTTPServerAddress)
	httpServer := rest.NewRestServer(handler, l)
	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		l.Info("error httpServer Notify", zap.String("Error: ", err.Error()))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error("app - Run - httpServer.Shutdown", zap.String("Error: ", err.Error()))
	}
}

func runTaskProcessor(ctx context.Context, cfg config.Config, redisOpt asynq.RedisClientOpt, q repo.Querier, l logger.Interface) {
	mailer := email.NewGmailSender(cfg.EmailSenderName, cfg.EmailSenderAddress, cfg.EmailSenderPassword)
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, mailer, q, l)
	fmt.Println("start task processor")
	err := taskProcessor.Start()
	if err != nil {
		fmt.Println("failed to start task processor")
	}
}
