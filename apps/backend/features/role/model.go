package role

import (
	"github.com/followthepattern/adapticc/models"
	"github.com/followthepattern/adapticc/types"
	validation "github.com/go-ozzo/ozzo-validation"
)

type RoleModel struct {
	ID   types.String `db:"id"`
	Code types.String `db:"code"`
	Name types.String `db:"name"`
	models.Userlog
}

func (m RoleModel) CreateValidate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Code, validation.Required),
		validation.Field(&m.Name, validation.Required),
	)
}

func (m RoleModel) UpdateValidate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.ID, validation.Required),
	)
}

func (m RoleModel) IsDefault() bool {
	return m.ID.Len() < 1
}

type RoleListRequestParams = models.ListRequestParams[models.ListFilter]

type RoleListResponse = models.ListResponse[RoleModel]

type UserRoleModel struct {
	UserID types.String `db:"user_id"`
	RoleID types.String `db:"role_id"`
	models.Userlog
}

func (m UserRoleModel) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.UserID, validation.Required),
		validation.Field(&m.RoleID, validation.Required),
	)
}
