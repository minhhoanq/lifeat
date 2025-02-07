package rest

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/hibiken/asynq"
	"github.com/minhhoanq/lifeat/common/logger"
	"github.com/minhhoanq/lifeat/user_service/config"
	v1 "github.com/minhhoanq/lifeat/user_service/internal/controller/rest/v1"
	"github.com/minhhoanq/lifeat/user_service/internal/controller/rest/v1/middleware"
	"github.com/minhhoanq/lifeat/user_service/internal/email"
	"github.com/minhhoanq/lifeat/user_service/internal/token"
	"github.com/minhhoanq/lifeat/user_service/internal/usecase"
	"github.com/minhhoanq/lifeat/user_service/internal/usecase/repo"
	"github.com/minhhoanq/lifeat/user_service/internal/worker"
	"github.com/minhhoanq/lifeat/user_service/pkg/constants"
	"github.com/minhhoanq/lifeat/user_service/pkg/postgres"
	"github.com/minhhoanq/lifeat/user_service/pkg/rest"
	"golang.org/x/sync/errgroup"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func RunRestServer(cfg config.Config) {
	l := logger.NewWrapLogger(zap.DebugLevel, false)

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	pg, err := postgres.New(cfg, l)
	if err != nil {
		l.Error("cannot connection to database", zap.String("Error: ", err.Error()))
		return
	}
	defer pg.Close(l)

	tokenMaker, err := token.NewJWTMaker(constants.PublicKeyPath, constants.PrivateKeyPath)
	if err != nil {
		l.Error("jwt maker error", zap.String("Error: ", err.Error()))
		return
	}

	redisOpts := asynq.RedisClientOpt{
		Addr: cfg.RedisAddres,
	}

	taskDistributor := worker.NewRedisTaskDistributor(l, redisOpts)

	db := postgres.Database{DB: pg.DB}
	q := repo.New(db)

	// Resful
	handler := echo.New()
	// CORS
	handler.Use(middleware.CORS)
	u := usecase.New(q, tokenMaker, cfg, taskDistributor)
	v1.NewRouter(handler, l, u, tokenMaker)

	// waitGroup
	waitGroup, ctx := errgroup.WithContext(ctx)

	// , rest.Port(cfg.HTTPServerAddress)
	// Waiting signal

	rest.NewRestServer(handler, l, waitGroup, ctx)
	runTaskProcessor(ctx, cfg, redisOpts, waitGroup, q, l)

	err = waitGroup.Wait()
	if err != nil {
		l.Error("error from wait group")
	}
}

func runTaskProcessor(ctx context.Context, cfg config.Config, redisOpt asynq.RedisClientOpt, waitGroup *errgroup.Group, q repo.Querier, l logger.Interface) {
	mailer := email.NewGmailSender(cfg.EmailSenderName, cfg.EmailSenderAddress, cfg.EmailSenderPassword)
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, mailer, q, l)
	l.Info("start task processor")
	err := taskProcessor.Start()
	if err != nil {
		l.Error("failed to start task processor", zap.String("ERROR", err.Error()))
	}

	waitGroup.Go(func() error {
		<-ctx.Done()
		l.Info("graceful shudown task processor")
		taskProcessor.Shutdown()
		l.Info("task processor stopped")
		return nil
	})
}
