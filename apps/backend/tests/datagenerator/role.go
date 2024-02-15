package datagenerator

import (
	"github.com/followthepattern/adapticc/features/role"
	"github.com/followthepattern/adapticc/types"
	"github.com/google/uuid"
)

func NewRandomRole() role.RoleModel {
	return role.RoleModel{
		ID:   types.StringFrom(uuid.NewString()),
		Name: types.StringFrom(String(8)),
	}
}
