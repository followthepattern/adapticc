package models

import "time"

type LoginResponse struct {
	JWT       *string    `json:"jwt,omitempty"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}
