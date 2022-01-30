package main

import (
	"backend/internal"
	"backend/internal/api"
	"backend/internal/config"
	"backend/internal/container"
	"backend/internal/hostserver"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	db, err := sql.Open("postgres", cfg.DB.ConnectionURL())

	if err != nil {
		log.Fatal(err)
	}

	var logger *zap.Logger

	if cfg.Api.Mode == config.ModeDev {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		log.Fatal(err)
	}

	cont := container.New(
		&ctx,
		cfg,
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

	if err := server.Serve(ctx, cfg.Api.Host, cfg.Api.Port); err != nil {
		logger.Error("failed to serve server", zap.Error(err))
	}
}
