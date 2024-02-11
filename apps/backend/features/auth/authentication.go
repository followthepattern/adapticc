package auth

import (
	"context"
	"crypto"
	"errors"
	"fmt"
	"time"

	"github.com/followthepattern/adapticc/config"
	"github.com/followthepattern/adapticc/features/mail"

	"github.com/followthepattern/adapticc/types"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const WRONG_EMAIL_OR_PASSWORD = "WRONG_EMAIL_OR_PASSWORD"
const EMAIL_IS_ALREADY_IN_USE_PATTERN = "%v is already in use, please try a different email address"

type AuthenticationService struct {
	repository AuthDatabase
	cfg        config.Config
	mail       mail.Mail
	jwtKeys    config.JwtKeyPair
}

func NewAuthenticationService(cfg config.Config, repository AuthDatabase, emailClient mail.Email, jwtKeys config.JwtKeyPair) AuthenticationService {
	return AuthenticationService{
		cfg:        cfg,
		repository: repository,
		mail:       mail.NewMail(cfg.Mail, emailClient),
		jwtKeys:    jwtKeys,
	}
}

func (service AuthenticationService) Login(ctx context.Context, email types.String, password types.String) (*LoginResponse, error) {
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

	tokenString, err := GenerateTokenStringFromUser(authUser, service.jwtKeys.Private)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		ExpiresAt: expiresAt,
		JWT:       tokenString,
	}, nil
}

func (service AuthenticationService) Register(ctx context.Context, register RegisterRequestParams) (*RegisterResponse, error) {
	ctxu := GetModelFromContext[AuthUser](ctx, CtxUserKey)
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

	creationUser := AuthUser{
		ID:        types.StringFrom(uuid.NewString()),
		Email:     register.Email,
		FirstName: register.FirstName,
		LastName:  register.LastName,
		Active:    types.FALSE,

		PasswordHash: string(passwordHash),
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

	return &RegisterResponse{
		Email:     creationUser.Email,
		FirstName: creationUser.FirstName,
		LastName:  creationUser.LastName,
	}, nil
}

func GenerateTokenStringFromUser(model AuthUser, privateKey crypto.PrivateKey) (string, error) {
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

func GetActivationMailTemplate(cfg config.Config, userID types.String, email types.String) mail.MailModel {
	activationLink := fmt.Sprintf("%s/users/activate/%s", cfg.Organization.Url, userID)

	from := fmt.Sprintf("%s <%s>", cfg.Organization.Name, cfg.Organization.Email)

	m := mail.MailModel{
		From:    from,
		To:      []string{email.Data},
		Subject: "Activate your email address",
		Text:    []byte(fmt.Sprintf("your activation link: %s", activationLink)),
	}

	return m
}
