package auth

import (
	"context"

	"github.com/followthepattern/adapticc/accesscontrol"
	"github.com/followthepattern/adapticc/container"
)

type AuthorizationService struct {
	repository AuthDatabase
	ac         accesscontrol.AccessControl
}

func NewAuthorizationService(cont container.Container, kind string) AuthorizationService {
	return AuthorizationService{
		ac: cont.GetAccessControl().WithKind(kind),
	}
}

func (service AuthorizationService) AuthorizedUser(ctx context.Context, action string, resourceID string) (string, error) {
	ctxu, err := GetUserContext(ctx)
	if err != nil {
		return "", err
	}

	roles, err := service.repository.GetRoleIDs(ctxu.ID.Data)
	if err != nil {
		return "", err
	}

	return ctxu.ID.Data, service.ac.Authorize(ctx, ctxu.ID.Data, action, resourceID, roles...)
}

func (service AuthorizationService) Authorize(ctx context.Context, action string, resourceID string) error {
	ctxu, err := GetUserContext(ctx)
	if err != nil {
		return err
	}

	roles, err := service.repository.GetRoleIDs(ctxu.ID.Data)
	if err != nil {
		return err
	}

	return service.ac.Authorize(ctx, ctxu.ID.Data, action, resourceID, roles...)
}
