package datagenerator

import (
	"github.com/followthepattern/adapticc/models"
	"github.com/followthepattern/adapticc/types"
	"github.com/google/uuid"
)

func NewRandomRole() models.Role {
	return models.Role{
		ID:   types.StringFrom(uuid.NewString()),
		Name: types.StringFrom(String(8)),
	}
}
