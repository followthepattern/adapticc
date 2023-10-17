package datagenerator

import (
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/google/uuid"
)

func NewRandomRole() models.Role {
	return models.Role{
		ID:   uuid.New().String(),
		Name: String(8),
	}
}
