package services

import (
	"context"
	"crypto"
	"errors"
	"fmt"
	"time"

	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/repositories/database"
	"github.com/followthepattern/adapticc/pkg/repositories/email"
	"github.com/followthepattern/adapticc/pkg/types"
	"github.com/followthepattern/adapticc/pkg/utils"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

const WRONG_EMAIL_OR_PASSWORD = "WRONG_EMAIL_OR_PASSWORD"
const EMAIL_IS_ALREADY_IN_USE_PATTERN = "%v is already in use, please try a different email address"

type Auth struct {
	repository database.Auth
	cfg        config.Config
	mail       Mail
	jwtKeys    config.JwtKeyPair
}

func NewAuth(cfg config.Config, repository database.Auth, emailClient email.Email, jwtKeys config.JwtKeyPair) Auth {
	return Auth{
		cfg:        cfg,
		repository: repository,
		mail:       NewMail(cfg.Mail, emailClient),
		jwtKeys:    jwtKeys,
	}
}

func (service Auth) Login(ctx context.Context, email types.String, password types.String) (*models.LoginResponse, error) {
	authUser, err := service.repository.VerifyLogin(email)
	if err != nil {
		return nil, err
	}

	if authUser.IsDefault() {
		return nil, errors.New(WRONG_EMAIL_OR_PASSWORD)
	}

	requestPasswordHash := utils.GeneratePasswordHash(password, authUser.Salt)

	if requestPasswordHash != authUser.PasswordHash.Data {
		return nil, errors.New(WRONG_EMAIL_OR_PASSWORD)
	}

	expiresAt := time.Now().Add(time.Hour * 24)

	tokenString, err := GenerateTokenStringFromUser(authUser.User, service.jwtKeys.Private)
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

	salt := types.StringFrom(utils.GenerateSaltString())
	passwordHash := types.StringFrom(utils.GeneratePasswordHash(register.Password, salt))

	creationUser := models.AuthUser{
		User: models.User{
			ID:        types.StringFrom(uuid.NewString()),
			Email:     register.Email,
			FirstName: register.FirstName,
			LastName:  register.LastName,
			Active:    types.FALSE,
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

	mail := GetActivationMailTemplate(service.cfg, creationUser.ID, creationUser.Email)

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

func GetActivationMailTemplate(cfg config.Config, userID types.String, email types.String) models.Mail {
	activationLink := fmt.Sprintf("%s/users/activate/%s", cfg.Organization.Url, userID)

	from := fmt.Sprintf("%s <%s>", cfg.Organization.Name, cfg.Organization.Email)

	m := models.Mail{
		From:    from,
		To:      []string{email.Data},
		Subject: "Activate your email address",
		Text:    []byte(fmt.Sprintf("your activation link: %s", activationLink)),
	}

	return m
}

func GenerateTokenStringFromUser(model models.User, privateKey crypto.PrivateKey) (string, error) {
	expiresAt := time.Now().Add(time.Hour * 24)
	token := jwt.NewWithClaims(jwt.GetSigningMethod("EdDSA"), jwt.MapClaims{
		"ID":        model.ID,
		"email":     model.Email,
		"firstName": model.FirstName,
		"lastName":  model.LastName,
		"expiresAt": expiresAt,
	})

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
