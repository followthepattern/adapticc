package datagenerator

import (
	"github.com/followthepattern/adapticc/features/auth"
	"github.com/followthepattern/adapticc/features/user"
	"github.com/followthepattern/adapticc/models"
	"github.com/followthepattern/adapticc/types"

	"github.com/google/uuid"
)

func NewRandomUser() user.UserModel {
	return user.UserModel{
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

func NewRandomAuthUser(password []byte) auth.AuthUser {
	return auth.AuthUser{
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
		PasswordHash: string(password),
	}
}
