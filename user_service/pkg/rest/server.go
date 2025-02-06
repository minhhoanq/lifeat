package rest

import (
	"context"
	"net/http"
	"time"

	"github.com/minhhoanq/lifeat/common/logger"
	"github.com/minhhoanq/lifeat/user_service/pkg/constants"

	"go.uber.org/zap"
)

const (
	_defaultShutdownTimeout = constants.DefaultShutdownTimeout
	_defaultAddr            = constants.DefaultPort
	_defaultReadTimeout     = constants.DefaultReadTimeout
	_defaultWriteTimeout    = constants.DefaultWriteTimeout
)

type RestServer struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

func NewRestServer(handler http.Handler, l logger.Interface, opts ...Option) *RestServer {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
		Addr:         _defaultAddr,
	}

	s := &RestServer{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
	}

	// Custom options
	for _, opt := range opts {
		opt(s)
	}

	s.Start(l)

	return s
}

func (s *RestServer) Start(l logger.Interface) {
	l.Info("start rest server at",
		zap.String("Adr: ", s.server.Addr),
		zap.String("ReadTimeout: ", s.server.ReadTimeout.String()),
		zap.String("WriteTimeout: ", s.server.WriteTimeout.String()),
		zap.String("ShutdownTimeout: ", s.shutdownTimeout.String()))
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *RestServer) Notify() <-chan error {
	return s.notify
}

func (s *RestServer) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
