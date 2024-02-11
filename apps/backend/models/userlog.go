package models

import "github.com/followthepattern/adapticc/types"

type Userlog struct {
	CreationUserID types.String `db:"creation_user_id" goqu:"skipupdate,omitempty"`
	UpdateUserID   types.String `db:"update_user_id" goqu:"omitempty"`
	CreatedAt      types.Time   `db:"created_at" goqu:"skipupdate,omitempty"`
	UpdatedAt      types.Time   `db:"updated_at" goqu:"omitempty"`
}
