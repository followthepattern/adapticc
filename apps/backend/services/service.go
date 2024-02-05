package services

import (
	"context"
	"errors"
	"log/slog"

	cerbos "github.com/cerbos/cerbos/client"
	"github.com/followthepattern/adapticc/config"
)

type Service struct {
	name         string
	logger       *slog.Logger
	cfg          config.Config
	cerbosClient cerbos.Client
}

func NewService(name string, cerbosClient cerbos.Client, cfg config.Config, logger *slog.Logger) Service {
	service := Service{
		name:         name,
		cerbosClient: cerbosClient,
		logger:       logger,
		cfg:          cfg,
	}

	return service
}

func (service Service) Authorize(ctx context.Context, action string, resourceID string, principalID string, roles ...string) error {
	principal := cerbos.NewPrincipal(principalID, roles...)

	resource := cerbos.NewResource(service.name, resourceID)

	ok, err := service.cerbosClient.IsAllowed(ctx, principal, resource, action)
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("action is not allowed")
	}

	return nil
}
