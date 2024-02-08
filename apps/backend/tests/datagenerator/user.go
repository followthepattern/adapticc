package datagenerator

import (
	"github.com/followthepattern/adapticc/models"
	"github.com/followthepattern/adapticc/types"

	"github.com/google/uuid"
)

func NewRandomUser() models.User {
	return models.User{
		ID:        types.StringFrom(uuid.NewString()),
		Email:     types.StringFrom(RandomEmail(8, 8)),
		FirstName: types.StringFrom(String(8)),
		LastName:  types.StringFrom(String(8)),
		Active:    types.FALSE,
		Userlog: models.Userlog{
			CreationUserID: types.StringFrom(uuid.NewString()),
			UpdateUserID:   types.StringFrom(uuid.NewString()),
			CreatedAt:      types.TimeNow(),
			UpdatedAt:      types.TimeNow(),
		},
	}
}

func NewRandomAuthUser(password []byte) models.AuthUser {
	return models.AuthUser{
		User: models.User{
			ID:        types.StringFrom(uuid.NewString()),
			Email:     types.StringFrom(RandomEmail(8, 8)),
			FirstName: types.StringFrom(String(8)),
			LastName:  types.StringFrom(String(8)),
			Active:    types.FALSE,
			Userlog: models.Userlog{
				CreationUserID: types.StringFrom(uuid.NewString()),
				UpdateUserID:   types.StringFrom(uuid.NewString()),
				CreatedAt:      types.TimeNow(),
				UpdatedAt:      types.TimeNow(),
			},
		},

		Password: models.Password{
			PasswordHash: string(password),
		},
	}
}
