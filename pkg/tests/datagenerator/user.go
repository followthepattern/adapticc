package datagenerator

import (
	"time"

	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/utils"
	"github.com/followthepattern/adapticc/pkg/utils/pointers"

	"github.com/google/uuid"
)

func NewRandomUser() models.User {
	return models.User{
		ID:        pointers.ToPtr(uuid.New().String()),
		Email:     pointers.ToPtr(RandomEmail(8, 8)),
		FirstName: pointers.ToPtr(String(8)),
		LastName:  pointers.ToPtr(String(8)),
	}
}

func NewRandomAuthUser(password string) models.AuthUser {
	salt := utils.GenerateSaltString()
	return models.AuthUser{
		User: models.User{
			ID:        pointers.ToPtr(uuid.New().String()),
			Email:     pointers.ToPtr(RandomEmail(8, 8)),
			FirstName: pointers.ToPtr(String(8)),
			LastName:  pointers.ToPtr(String(8)),
			Active:    pointers.ToPtr(false),
			Userlog: models.Userlog{
				CreationUserID: pointers.ToPtr(uuid.New().String()),
				UpdateUserID:   pointers.ToPtr(uuid.New().String()),
				CreatedAt:      pointers.ToPtr(time.Now()),
				UpdatedAt:      pointers.ToPtr(time.Now()),
			},
		},
		Password: models.Password{
			Salt:         salt,
			PasswordHash: utils.GeneratePasswordHash(password, salt),
		},
	}
}
