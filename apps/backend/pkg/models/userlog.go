package models

import "time"

type Userlog struct {
	CreationUserID *string    `db:"creation_user_id" goqu:"skipupdate"`
	UpdateUserID   *string    `db:"update_user_id"`
	CreatedAt      *time.Time `db:"created_at" goqu:"skipupdate"`
	UpdatedAt      *time.Time `db:"updated_at"`
}
