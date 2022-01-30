package models

import "time"

type Userlog struct {
	CreationUserID *string    `json:"creation_user_id,omitempty" db:"creation_user_id"`
	UpdateUserID   *string    `json:"update_user_id,omitempty" db:"update_user_id"`
	CreatedAt      *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
