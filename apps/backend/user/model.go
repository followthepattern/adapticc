package user

import (
	"github.com/followthepattern/adapticc/models"
	"github.com/followthepattern/adapticc/types"
	validation "github.com/go-ozzo/ozzo-validation"
)

type UserModel struct {
	ID        types.String  `db:"id" goqu:"skipupdate,omitempty"`
	Email     types.String  `db:"email" goqu:"skipupdate,omitempty"`
	FirstName types.String  `db:"first_name" goqu:"omitempty"`
	LastName  types.String  `db:"last_name" goqu:"omitempty"`
	Active    types.Bool    `db:"active" goqu:"skipupdate,omitempty"`
	Roles     []models.Role `db:"-"`
	models.Userlog
}

func (u UserModel) CreateValidate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required),
		validation.Field(&u.FirstName, validation.Required),
		validation.Field(&u.LastName, validation.Required),
	)
}

func (u UserModel) UpdateValidate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.ID, validation.Required),
	)
}

func (u UserModel) IsDefault() bool {
	return u.ID.Len() < 1
}

type UserListRequestParams = models.ListRequestParams[models.ListFilter]

type UserListResponse = models.ListResponse[UserModel]
