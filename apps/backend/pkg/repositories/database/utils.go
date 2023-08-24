package database

import (
	"time"

	"github.com/followthepattern/adapticc/pkg/models"
)

func setCreateUserlog(userID string, createdAt time.Time) models.Userlog {
	var createUser *string

	if len(userID) > 0 {
		createUser = &userID
	}

	return models.Userlog{
		CreationUserID: createUser,
		CreatedAt:      &createdAt,
	}
}

func setUpdateUserlog(userID string, updatedAt time.Time) models.Userlog {
	return models.Userlog{
		UpdateUserID: &userID,
		UpdatedAt:    &updatedAt,
	}
}
