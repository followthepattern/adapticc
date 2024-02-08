package hostserver

import (
	"context"
	"fmt"
	"net/http"

	"log/slog"

	"github.com/oklog/run"
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

	var g run.Group

	g.Add(func() error {
		s.logger.Info(fmt.Sprintf("Server start listnening on %s:%s", host, port))
		return srv.ListenAndServe()
	}, func(error) {
		srv.Close()
		srv.Shutdown(ctx)
	})

	return g.Run()
}
