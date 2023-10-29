package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/repositories/database"
	"github.com/followthepattern/adapticc/pkg/utils"
	"github.com/followthepattern/adapticc/pkg/utils/pointers"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

const WRONG_EMAIL_OR_PASSWORD = "WRONG_EMAIL_OR_PASSWORD"
const EMAIL_IS_ALREADY_IN_USE_PATTERN = "%v is already in use, please try a different email address"

type Auth struct {
	repository database.Auth
	cfg        config.Config
	mail       Mail
}

func NewAuth(cfg config.Config, repository database.Auth) Auth {
	return Auth{
		cfg:        cfg,
		repository: repository,
		mail:       NewMail(cfg.Mail),
	}
}

func (service Auth) Login(ctx context.Context, email string, password string) (*models.LoginResponse, error) {
	authUser, err := service.repository.VerifyLogin(email)
	if err != nil {
		return nil, err
	}

	if authUser.IsDefault() {
		return nil, errors.New(WRONG_EMAIL_OR_PASSWORD)
	}

	requestPasswordHash := utils.GeneratePasswordHash(password, authUser.Salt)

	if requestPasswordHash != authUser.PasswordHash {
		return nil, errors.New(WRONG_EMAIL_OR_PASSWORD)
	}

	expiresAt := time.Now().Add(time.Hour * 24)
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"ID":        *authUser.ID,
		"email":     *authUser.Email,
		"firstName": *authUser.FirstName,
		"lastName":  *authUser.LastName,
		"expiresAt": expiresAt,
	})

	tokenString, err := token.SignedString([]byte(service.cfg.Server.HmacSecret))
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		ExpiresAt: expiresAt,
		JWT:       tokenString,
	}, nil
}

func (service Auth) Register(ctx context.Context, register models.RegisterRequestParams) (*models.RegisterResponse, error) {
	ctxu := utils.GetModelFromContext[models.User](ctx, utils.CtxUserKey)
	if ctxu == nil {
		return nil, fmt.Errorf("invalid user context")
	}

	valid, err := service.repository.VerifyEmail(register.Email)
	if err != nil {
		return nil, err
	}

	if !valid {
		return nil, fmt.Errorf(EMAIL_IS_ALREADY_IN_USE_PATTERN, register.Email)
	}

	salt := utils.GenerateSaltString()
	passwordHash := utils.GeneratePasswordHash(register.Password, salt)

	creationUser := models.AuthUser{
		User: models.User{
			ID:        pointers.ToPtr(uuid.New().String()),
			Email:     &register.Email,
			FirstName: &register.FirstName,
			LastName:  &register.LastName,
			Active:    pointers.ToPtr(false),
		},
		Password: models.Password{
			PasswordHash: passwordHash,
			Salt:         salt,
		},
	}

	err = service.repository.RegisterUser(creationUser)
	if err != nil {
		return nil, err
	}

	mail := service.getActivationMail(*creationUser.ID, *creationUser.Email)

	err = service.mail.SendMail(mail)
	if err != nil {
		return nil, err
	}

	return &models.RegisterResponse{
		Email:     creationUser.Email,
		FirstName: creationUser.FirstName,
		LastName:  creationUser.LastName,
	}, nil
}

func (service Auth) getActivationMail(userID string, email string) models.Mail {
	activationLink := fmt.Sprintf("%s/users/activate/%s", service.cfg.Organization.Url, userID)

	from := fmt.Sprintf("%s <%s>", service.cfg.Organization.Name, service.cfg.Organization.Email)

	m := models.Mail{
		From:    from,
		To:      []string{email},
		Subject: "Activate your email address",
		Text:    []byte(fmt.Sprintf("your activation link: %s", activationLink)),
	}

	return m
}

func GenerateTokenStringFromUser(model models.User, secret []byte) (string, error) {
	expiresAt := time.Now().Add(time.Hour * 24)
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"ID":        *model.ID,
		"email":     *model.Email,
		"firstName": *model.FirstName,
		"lastName":  *model.LastName,
		"expiresAt": expiresAt,
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
