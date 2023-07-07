package middlewares

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
)

type LoggerInterface interface {
	Print(v ...interface{})
}

type logPrinter func(v ...interface{})

func (f logPrinter) Print(v ...interface{}) {
	f(v)
}

func AddMiddlewareLogger(r *chi.Mux, logger *zap.Logger) {
	logFunc := logPrinter(func(values ...interface{}) {
		logger.Info("HTTP", zap.Any("values", values))
	})

	middleware.DefaultLogger = middleware.RequestLogger(&middleware.DefaultLogFormatter{
		Logger:  logFunc,
		NoColor: true,
	})

	r.Use(middleware.Logger)
}
