package models

import "time"

type Userlog struct {
	CreationUserID string    `db:"creation_user_id" goqu:"skipupdate,omitempty"`
	UpdateUserID   string    `db:"update_user_id" goqu:"omitempty"`
	CreatedAt      time.Time `db:"created_at" goqu:"skipupdate,omitempty"`
	UpdatedAt      time.Time `db:"updated_at" goqu:"omitempty"`
}
