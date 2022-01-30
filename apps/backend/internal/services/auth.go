package services

import (
	"backend/internal/container"
	"backend/internal/models"
	repositories "backend/internal/repositories/database"
	"backend/internal/utils"
	"backend/internal/utils/pointers"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

const INCORRECT_EMAIL_OR_PASSWORD = "incorrect email address or password, please try again"
const EMAIL_IS_ALREADY_IN_USE_PATTERN = "%v is already in use, please try a different email address"

type Auth struct {
	userRepository      repositories.User
	userTokenRepository repositories.UserToken
}

func AuthDependencyConstructor(cont container.IContainer) (interface{}, error) {
	dependency := Auth{}

	ur, err := repositories.ResolveUser(cont)
	if err != nil {
		return nil, err
	}

	urt, err := repositories.ResolveUserToken(cont)
	if err != nil {
		return nil, err
	}

	dependency.userRepository = *ur
	dependency.userTokenRepository = *urt

	return &dependency, nil
}

func (a Auth) Login(email string, password string) (*models.LoginResponse, error) {
	user, err := a.userRepository.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	if user == nil || user.ID == nil {
		return nil, errors.New(INCORRECT_EMAIL_OR_PASSWORD)
	}

	requestPasswordHash := utils.GeneratePasswordHash(password, *user.Salt)

	if requestPasswordHash != *user.PasswordHash {
		return nil, errors.New(INCORRECT_EMAIL_OR_PASSWORD)
	}
	jwt := utils.GenerateSaltString()

	token := models.UserToken{
		UserID:    user.ID,
		Token:     &jwt,
		ExpiresAt: pointers.Time(time.Now()),
	}

	err = a.userTokenRepository.Create(&token)
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		JWT:       pointers.String(jwt),
		ExpiresAt: token.ExpiresAt,
	}, nil
}

func (a Auth) Register(email string, firstName string, lastName string, password string) (*models.RegisterResponse, error) {
	if user, err := a.userRepository.GetByEmail(email); err != nil {
		return nil, err
	} else if !user.IsNil() {
		return nil, fmt.Errorf(EMAIL_IS_ALREADY_IN_USE_PATTERN, *user.Email)
	}

	salt := utils.GenerateSaltString()
	passwordHash := utils.GeneratePasswordHash(password, salt)
	log := models.Userlog{
		CreatedAt: pointers.Time(time.Now()),
	}

	user := models.User{
		ID:           pointers.String(uuid.New().String()),
		Email:        &email,
		FirstName:    &firstName,
		LastName:     &lastName,
		PasswordHash: &passwordHash,
		Salt:         &salt,
		Active:       pointers.Bool(true),
		Userlog:      log,
	}

	if err := a.userRepository.Create(&user); err != nil {
		return nil, err
	}

	return &models.RegisterResponse{
		Email:     &email,
		FirstName: &firstName,
		LastName:  &lastName,
	}, nil
}
