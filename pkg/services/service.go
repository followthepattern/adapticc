package services

import (
	"context"
	"errors"

	cerbos "github.com/cerbos/cerbos/client"
	"github.com/followthepattern/adapticc/pkg/config"
	"go.uber.org/zap"
)

type getRoles interface {
	GetRolesByUserID(userID string) ([]string, error)
}

type Service struct {
	getRoles
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

	service.getRoles = service

	return service
}

func (service Service) Authorize(ctx context.Context, principalID string, action string, resourceID string) error {
	roles, err := service.getRoles.GetRolesByUserID(principalID)
	if err != nil {
		return err
	}

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

func (service Service) GetRolesByUserID(userID string) ([]string, error) {
	return []string{}, nil
}
