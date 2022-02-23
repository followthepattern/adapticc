package user

import (
	"backend/internal/models"
	"backend/internal/utils"

	"github.com/graph-gophers/graphql-go"
)

type user struct {
	ID             string
	Email          string
	FirstName      string
	LastName       string
	Active         *bool
	LastLoginAt    *graphql.Time
	CreatedAt      graphql.Time
	CreationUserID *string
	UpdatedAt      *graphql.Time
	UpdateUserID   *string
}

func GetFromModel(model *models.User) *user {
	if model.IsNil() {
		return nil
	}

	result := user{
		ID:             *model.ID,
		Email:          *model.Email,
		FirstName:      *model.FirstName,
		LastName:       *model.LastName,
		Active:         model.Active,
		LastLoginAt:    utils.TimeToGraphqlTime(model.LastLoginAt),
		CreatedAt:      *utils.TimeToGraphqlTime(model.CreatedAt),
		CreationUserID: model.CreationUserID,
		UpdatedAt:      utils.TimeToGraphqlTime(model.UpdatedAt),
		UpdateUserID:   model.UpdateUserID,
	}

	return &result
}
