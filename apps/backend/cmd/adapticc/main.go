package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/followthepattern/adapticc/pkg/accesscontrol"
	"github.com/followthepattern/adapticc/pkg/api"
	"github.com/followthepattern/adapticc/pkg/api/graphql"
	"github.com/followthepattern/adapticc/pkg/api/rest"
	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/controllers"
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

	cerbosClient, err := accesscontrol.New(cfg.Cerbos)
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctrls := controllers.New(ctx, cerbosClient, db, *cfg, logger)

	graphqlHandler := graphql.New(ctrls)

	restHandler := rest.New(ctrls)

	router := api.NewHttpApi(*cfg, graphqlHandler, restHandler, logger)

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
