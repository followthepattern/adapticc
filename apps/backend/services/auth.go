package services

import (
	"context"
	"crypto"
	"errors"
	"fmt"
	"time"

	"github.com/followthepattern/adapticc/config"
	"github.com/followthepattern/adapticc/models"
	"github.com/followthepattern/adapticc/repositories/database"
	"github.com/followthepattern/adapticc/repositories/email"
	"github.com/followthepattern/adapticc/types"
	"github.com/followthepattern/adapticc/user"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const WRONG_EMAIL_OR_PASSWORD = "WRONG_EMAIL_OR_PASSWORD"
const EMAIL_IS_ALREADY_IN_USE_PATTERN = "%v is already in use, please try a different email address"

type ContextKey struct {
	Name string
}

var CtxUserKey = ContextKey{Name: "ctx-user"}

func GetModelFromContext[T any](ctx context.Context, ctxKey ContextKey) *T {
	obj := ctx.Value(ctxKey)

	model, ok := obj.(T)
	if !ok {
		return nil
	}

	return &model
}

func GetUserContext(ctx context.Context) (user.UserModel, error) {
	obj := ctx.Value(CtxUserKey)

	model, ok := obj.(user.UserModel)
	if !ok {
		return user.UserModel{}, errors.New("invalid user context")
	}

	if model.IsDefault() {
		return user.UserModel{}, errors.New("invalid user context")
	}

	return model, nil
}

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

	if err = bcrypt.CompareHashAndPassword([]byte(authUser.PasswordHash), []byte(password.Data)); err != nil {
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
	ctxu := GetModelFromContext[user.UserModel](ctx, CtxUserKey)
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

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(register.Password.Data), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	creationUser := models.AuthUser{
		User: user.UserModel{
			ID:        types.StringFrom(uuid.NewString()),
			Email:     register.Email,
			FirstName: register.FirstName,
			LastName:  register.LastName,
			Active:    types.FALSE,
		},
		Password: models.Password{
			PasswordHash: string(passwordHash),
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
