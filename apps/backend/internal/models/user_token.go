package models

import "time"

type UserToken struct {
	UserID    *string    `json:"user_id" db:"user_id"`
	Token     *string    `json:"token" db:"token"`
	ExpiresAt *time.Time `json:"expires_at" db:"expires_at"`
}
