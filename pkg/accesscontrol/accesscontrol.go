package accesscontrol

import (
	"context"
	"errors"

	cerbos "github.com/cerbos/cerbos/client"
	"github.com/followthepattern/adapticc/pkg/config"
)

const (
	CREATE = "create"
	READ   = "read"
	UPDATE = "update"
	DELETE = "delete"

	ALLRESOURCE = "ALL"
)

type Config struct {
	Kind   string
	Cerbos cerbos.Client
}

func (c Config) Build() AccessControl {
	return AccessControl{
		kind:   c.Kind,
		cerbos: c.Cerbos,
	}
}

type AccessControl struct {
	kind   string
	cerbos cerbos.Client
}

func NewClient(cfg config.Cerbos) (cerbos.Client, error) {
	client, err := cerbos.New(cfg.Address, cerbos.WithPlaintext())
	if err != nil {
		return nil, err
	}

	return client, nil
}

func New(cfg config.Cerbos) (AccessControl, error) {
	client, err := cerbos.New(cfg.Address, cerbos.WithPlaintext())
	if err != nil {
		return AccessControl{}, err
	}

	return AccessControl{
		cerbos: client,
	}, nil
}

func (ac AccessControl) WithKind(kind string) AccessControl {
	ac.kind = kind
	return ac
}

func (ac AccessControl) Authorize(ctx context.Context, principalID string, action string, resourceID string, roles ...string) error {
	principal := cerbos.NewPrincipal(principalID, roles...)

	resource := cerbos.NewResource(ac.kind, resourceID)

	ok, err := ac.cerbos.IsAllowed(ctx, principal, resource, action)
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("action is not allowed")
	}

	return nil
}
