package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"

	"log/slog"

	"github.com/followthepattern/adapticc/pkg/accesscontrol"
	"github.com/followthepattern/adapticc/pkg/api"
	"github.com/followthepattern/adapticc/pkg/api/graphql"
	"github.com/followthepattern/adapticc/pkg/api/graphql/schema"
	"github.com/followthepattern/adapticc/pkg/api/rest"
	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/followthepattern/adapticc/pkg/controllers"
	"github.com/followthepattern/adapticc/pkg/hostserver"
	"github.com/followthepattern/adapticc/pkg/repositories/email"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}

	jwt, err := config.ReadKeys(cfg.Server)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.DB.ConnectionURL())
	if err != nil {
		log.Fatal(err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	cerbosClient, err := accesscontrol.New(cfg.Cerbos)
	if err != nil {
		log.Fatal(err)
	}

	schemaDef, err := schema.GetSchema(cfg.Server)
	if err != nil {
		log.Fatal(err)
	}

	emailClient := email.NewClient()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	cont := container.New(cerbosClient, emailClient, db, *cfg, logger, jwt)

	ctrls := controllers.New(cont)

	graphqlHandler := graphql.New(ctrls, schemaDef)

	restHandler := rest.New(ctrls)

	router := api.NewHttpApi(*cfg, jwt, graphqlHandler, restHandler, logger)

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
