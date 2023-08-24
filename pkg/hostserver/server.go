package hostserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

const serveShutDownTimeout = 5

type Server struct {
	logger *zap.Logger
	router http.Handler
}

func NewServer(router http.Handler, logger *zap.Logger) *Server {
	return &Server{
		logger: logger,
		router: router,
	}
}

func (s Server) Serve(ctx context.Context, host string, port string) (err error) {
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: s.router,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatal("server listen err:", zap.Error(err))
		}
	}()

	s.logger.Info(fmt.Sprintf("Server start listnening on %s:%s", host, port))

	<-ctx.Done()

	s.logger.Info("Server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), serveShutDownTimeout*time.Second)
	defer cancel()

	if err = srv.Shutdown(ctxShutDown); err != nil {
		s.logger.Fatal("Server Shutdown Failed: ", zap.Error(err))
	}
	s.logger.Info("Server Graceful shutdown success")

	if err == http.ErrServerClosed {
		err = nil
	}

	return nil
}
