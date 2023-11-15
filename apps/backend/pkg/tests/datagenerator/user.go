package datagenerator

import (
	"time"

	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/types"
	"github.com/followthepattern/adapticc/pkg/utils"

	"github.com/google/uuid"
)

func NewRandomUser() models.User {
	return models.User{
		ID:        types.StringFrom(uuid.NewString()),
		Email:     types.StringFrom(RandomEmail(8, 8)),
		FirstName: types.StringFrom(String(8)),
		LastName:  types.StringFrom(String(8)),
		Active:    false,
		Userlog: models.Userlog{
			CreationUserID: types.StringFrom(uuid.NewString()),
			UpdateUserID:   types.StringFrom(uuid.NewString()),
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
	}
}

func NewRandomAuthUser(password string) models.AuthUser {
	salt := types.StringFrom(utils.GenerateSaltString())
	passwordString := types.StringFrom(password)
	return models.AuthUser{
		User: models.User{
			ID:        types.StringFrom(uuid.NewString()),
			Email:     types.StringFrom(RandomEmail(8, 8)),
			FirstName: types.StringFrom(String(8)),
			LastName:  types.StringFrom(String(8)),
			Active:    false,
			Userlog: models.Userlog{
				CreationUserID: types.StringFrom(uuid.NewString()),
				UpdateUserID:   types.StringFrom(uuid.NewString()),
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
		},
		Password: models.Password{
			Salt:         salt,
			PasswordHash: types.StringFrom(utils.GeneratePasswordHash(passwordString, salt)),
		},
	}
}
