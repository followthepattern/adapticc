package middlewares

import (
	"log/slog"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type LoggerInterface interface {
	Print(v ...interface{})
}

type logPrinter func(v ...interface{})

func (f logPrinter) Print(v ...interface{}) {
	f(v)
}

func AddMiddlewareLogger(r *chi.Mux, logger *slog.Logger) {
	logFunc := logPrinter(func(values ...interface{}) {
		logger.Debug("HTTP", slog.Any("values", values))
	})

	middleware.DefaultLogger = middleware.RequestLogger(&middleware.DefaultLogFormatter{
		Logger:  logFunc,
		NoColor: true,
	})

	r.Use(middleware.Logger)
}
