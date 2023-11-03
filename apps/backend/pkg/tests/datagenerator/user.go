package datagenerator

import (
	"time"

	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/utils"

	"github.com/google/uuid"
)

func NewRandomUser() models.User {
	return models.User{
		ID:        uuid.New().String(),
		Email:     RandomEmail(8, 8),
		FirstName: String(8),
		LastName:  String(8),
		Active:    false,
		Userlog: models.Userlog{
			CreationUserID: uuid.New().String(),
			UpdateUserID:   uuid.New().String(),
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
	}
}

func NewRandomAuthUser(password string) models.AuthUser {
	salt := utils.GenerateSaltString()
	return models.AuthUser{
		User: models.User{
			ID:        uuid.New().String(),
			Email:     RandomEmail(8, 8),
			FirstName: String(8),
			LastName:  String(8),
			Active:    false,
			Userlog: models.Userlog{
				CreationUserID: uuid.New().String(),
				UpdateUserID:   uuid.New().String(),
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
		},
		Password: models.Password{
			Salt:         salt,
			PasswordHash: utils.GeneratePasswordHash(password, salt),
		},
	}
}
