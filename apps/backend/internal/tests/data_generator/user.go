package data_generator

import (
	"backend/internal/models"
	"backend/internal/utils"
	"backend/internal/utils/pointers"

	"github.com/google/uuid"
)

func NewRandomUser(password string) models.User {
	salt := utils.GenerateSaltString()
	return models.User{
		ID:           pointers.String(uuid.New().String()),
		Email:        pointers.String(String(15)),
		FirstName:    pointers.String(String(8)),
		LastName:     pointers.String(String(8)),
		Salt:         pointers.String(salt),
		PasswordHash: pointers.String(utils.GeneratePasswordHash(password, salt)),
	}
}
