package models

import (
	"time"

	"github.com/followthepattern/adapticc/pkg/types"
)

type Userlog struct {
	CreationUserID types.String `db:"creation_user_id" goqu:"skipupdate,omitempty"`
	UpdateUserID   types.String `db:"update_user_id" goqu:"omitempty"`
	CreatedAt      time.Time    `db:"created_at" goqu:"skipupdate,omitempty"`
	UpdatedAt      time.Time    `db:"updated_at" goqu:"omitempty"`
}
