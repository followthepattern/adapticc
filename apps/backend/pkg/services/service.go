package services

import (
	"context"
	"errors"

	cerbos "github.com/cerbos/cerbos/client"
	"github.com/followthepattern/adapticc/pkg/config"
	"go.uber.org/zap"
)

type Service struct {
	name         string
	logger       *zap.Logger
	cfg          config.Config
	cerbosClient cerbos.Client
}

func NewService(name string, cerbosClient cerbos.Client, cfg config.Config, logger *zap.Logger) Service {
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
