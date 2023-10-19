package hostserver

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"log/slog"
)

const serveShutDownTimeout = 5

type Server struct {
	logger *slog.Logger
	router http.Handler
}

func NewServer(router http.Handler, logger *slog.Logger) *Server {
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
			s.logger.Error("server listen err:", err)
			os.Exit(1)
		}
	}()

	s.logger.Info(fmt.Sprintf("Server start listnening on %s:%s", host, port))

	<-ctx.Done()

	s.logger.Info("Server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), serveShutDownTimeout*time.Second)
	defer cancel()

	if err = srv.Shutdown(ctxShutDown); err != nil {
		s.logger.Error("Server Shutdown Failed: ", err)
	}
	s.logger.Info("Server Graceful shutdown success")

	if err == http.ErrServerClosed {
		err = nil
	}

	return nil
}
