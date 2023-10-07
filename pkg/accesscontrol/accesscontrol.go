package accesscontrol

import (
	cerbos "github.com/cerbos/cerbos/client"
	"github.com/followthepattern/adapticc/pkg/config"
)

const (
	CREATE = "create"
	READ   = "read"
	UPDATE = "update"
	DELETE = "delete"
)

func NewClient(cfg config.Cerbos) (cerbos.Client, error) {
	client, err := cerbos.New(cfg.Address, cerbos.WithPlaintext())
	if err != nil {
		return nil, err
	}

	return client, nil
}
