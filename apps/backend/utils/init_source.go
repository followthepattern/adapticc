package utils

import (
	"context"
	"time"

	"log/slog"
)

const (
	defaultInitInterval = 5 * time.Second
	defaultInitTimeout  = 300
)

type InitFunc = func() error

func InitSource(ctx context.Context, initFunc InitFunc, source string, logger *slog.Logger) {
	isConnDB := func() bool {
		if err := initFunc(); err != nil {
			logger.Error("unable to init source:", err)
			return false
		}
		return true
	}

	ok := Retry(ctx, defaultInitInterval,
		time.Duration(defaultInitTimeout)*time.Second,
		isConnDB,
	)

	if !ok {
		logger.Error("Can't init source", slog.String("source", source))
	}
}
