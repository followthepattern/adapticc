package utils

import (
	"context"
	"time"

	"go.uber.org/zap"
)

const (
	defaultInitInterval = 5 * time.Second
	defaultInitTimeout  = 300
)

type InitFunc = func() error

func InitSource(ctx context.Context, initFunc InitFunc, source string, logger *zap.Logger) {

	isConnDB := func() bool {
		if err := initFunc(); err != nil {
			logger.Info("unable to init source:", zap.Error(err))
			return false
		}
		return true
	}

	ok := Retry(ctx, defaultInitInterval,
		time.Duration(defaultInitTimeout)*time.Second,
		isConnDB,
	)

	if !ok {
		logger.Fatal("Can't init source", zap.String("source", source))
	}
}
