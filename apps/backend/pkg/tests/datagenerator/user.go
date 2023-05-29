package datagenerator

import (
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/utils"
	"github.com/followthepattern/adapticc/pkg/utils/pointers"

	"github.com/google/uuid"
)

func NewRandomUser() models.User {
	return models.User{
		ID:        pointers.String(uuid.New().String()),
		Email:     pointers.String(RandomEmail(8, 8)),
		FirstName: pointers.String(String(8)),
		LastName:  pointers.String(String(8)),
	}
}

func NewRandomAuthUser(password string) models.User {
	salt := utils.GenerateSaltString()
	return models.User{
		ID:        pointers.String(uuid.New().String()),
		Email:     pointers.String(RandomEmail(8, 8)),
		FirstName: pointers.String(String(8)),
		LastName:  pointers.String(String(8)),
		Salt:      pointers.String(salt),
		Password:  pointers.String(utils.GeneratePasswordHash(password, salt)),
	}
}
