package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/followthepattern/adapticc/pkg/api"
	"github.com/followthepattern/adapticc/pkg/api/graphql"
	"github.com/followthepattern/adapticc/pkg/api/graphql/resolvers"
	"github.com/followthepattern/adapticc/pkg/api/rest"
	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/controllers"
	"github.com/followthepattern/adapticc/pkg/hostserver"
	"github.com/followthepattern/adapticc/pkg/repositories/database"
	"github.com/followthepattern/adapticc/pkg/services"

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

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	userController, err := buildUserController(ctx, db, *cfg, logger)
	if err != nil {
		log.Fatal(err)
	}

	authController, err := buildAuthController(ctx, db, *cfg, logger)
	if err != nil {
		log.Fatal(err)
	}

	productController, err := buildProductController(ctx, db, *cfg, logger)
	if err != nil {
		log.Fatal(err)
	}

	resolverConfig := resolvers.NewResolverConfig(
		*userController,
		*authController,
		*productController)

	resolver := resolvers.New(resolverConfig)

	graphqlHandler := graphql.New(&resolver)

	restHandler := rest.New(rest.NewRestConfig(*productController))

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

func buildUserController(ctx context.Context, db *sql.DB, cfg config.Config, logger *zap.Logger) (*controllers.User, error) {
	user, err := database.NewUser(ctx, db)
	if err != nil {
		return nil, err
	}
	ctrl := controllers.NewUser(user)

	return &ctrl, err
}

func buildAuthController(ctx context.Context, db *sql.DB, cfg config.Config, logger *zap.Logger) (*controllers.Auth, error) {
	auth, err := database.NewAuth(ctx, db, logger)
	if err != nil {
		return nil, err
	}
	authService := services.NewAuth(cfg, auth)

	ctrl := controllers.NewAuth(cfg, authService)

	return &ctrl, err
}

func buildProductController(ctx context.Context, db *sql.DB, cfg config.Config, logger *zap.Logger) (*controllers.Product, error) {
	product, err := database.NewProduct(ctx, db)
	if err != nil {
		return nil, err
	}
	productController := controllers.NewProduct(product)

	return &productController, err
}
