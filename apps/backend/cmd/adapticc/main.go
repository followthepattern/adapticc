package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"

	internal "github.com/followthepattern/adapticc/pkg"
	"github.com/followthepattern/adapticc/pkg/api"
	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/followthepattern/adapticc/pkg/hostserver"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.DB.ConnectionURL())
	if err != nil {
		log.Fatal(err)
	}

	logConfig := zap.NewProductionConfig()
	logConfig.Level = zap.NewAtomicLevelAt(zapcore.Level(cfg.Server.LogLevel))

	logger, err := logConfig.Build()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	cont := container.New(
		ctx,
		*cfg,
		db,
		logger)

	err = internal.RegisterDependencies(cont)
	if err != nil {
		log.Fatal(err.Error())
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	router, err := api.GetRouter(cont)
	if err != nil {
		log.Fatal(err)
	}

	server := hostserver.NewServer(router, logger)
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		osCall := <-c
		logger.Info(fmt.Sprintf("Stop server system call:%+v", osCall))
		cancel()
	}()

	if err := server.Serve(ctx, cfg.Server.Host, cfg.Server.Port); err != nil {
		logger.Error("failed to serve server", zap.Error(err))
	}
}
