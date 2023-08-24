package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/followthepattern/adapticc/pkg/models"
	repositories "github.com/followthepattern/adapticc/pkg/repositories/database"
	"github.com/followthepattern/adapticc/pkg/request"
	"github.com/followthepattern/adapticc/pkg/utils"
	"github.com/followthepattern/adapticc/pkg/utils/pointers"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

const WRONG_EMAIL_OR_PASSWORD = "WRONG_EMAIL_OR_PASSWORD"
const EMAIL_IS_ALREADY_IN_USE_PATTERN = "%v is already in use, please try a different email address"

type Auth struct {
	authMsgChannelIn  <-chan models.AuthMsg
	authMsgChannelOut chan<- models.AuthMsg
	ctx               context.Context
	cfg               config.Config
	sendMsg           func(ctx context.Context, msg models.AuthMsg) error
}

type AuthMsgChannel chan models.AuthMsg

func RegisterAuthChannel(cont *container.Container) {
	if cont == nil {
		return
	}
	authMsgChannel := make(AuthMsgChannel)
	container.Register(cont, func(cont *container.Container) (*AuthMsgChannel, error) {
		return &authMsgChannel, nil
	})
}

func AuthDependencyConstructor(cont *container.Container) (*Auth, error) {
	authMsgChannelIn, err := container.Resolve[AuthMsgChannel](cont)
	if err != nil {
		return nil, err
	}

	authMsgChannelOut, err := container.Resolve[repositories.AuthMsgChannel](cont)
	if err != nil {
		return nil, err
	}

	dependency := Auth{
		ctx:               cont.GetContext(),
		cfg:               cont.GetConfig(),
		authMsgChannelIn:  *authMsgChannelIn,
		authMsgChannelOut: *authMsgChannelOut,
		sendMsg:           request.CreateSenderFunc(*authMsgChannelOut, request.DefaultTimeOut),
	}

	go func() {
		dependency.MonitorChannels()
	}()

	return &dependency, nil
}

func (service Auth) MonitorChannels() {
	for {
		select {
		case request := <-service.authMsgChannelIn:
			if request.Login != nil {
				result, err := service.Login(request.Login.Context(), request.Login.RequestParams().Email, request.Login.RequestParams().Password)
				if err != nil {
					request.Login.ReplyError(err)
					continue
				}
				request.Login.Reply(*result)
			} else if request.Register != nil {
				register := request.Register.RequestParams()
				result, err := service.Register(request.Register.Context(), register)
				if err != nil {
					request.Register.ReplyError(err)
					continue
				}
				request.Register.Reply(*result)
			}
		case <-service.ctx.Done():
			return
		}
	}
}

func (a Auth) Login(ctx context.Context, email string, password string) (*models.LoginResponse, error) {
	req := request.New[string, models.AuthUser](ctx, email)

	msg := models.AuthMsg{
		VerifyLogin: &req,
	}

	if err := a.sendMsg(ctx, msg); err != nil {
		return nil, err
	}

	authUser, err := req.Wait()
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

	tokenString, err := token.SignedString([]byte(a.cfg.Server.HmacSecret))
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		ExpiresAt: expiresAt,
		JWT:       tokenString,
	}, nil
}

func (a Auth) Register(ctx context.Context, register models.RegisterRequestParams) (*models.RegisterResponse, error) {
	ctxu := utils.GetModelFromContext[models.User](ctx, utils.CtxUserKey)
	if ctxu == nil {
		return nil, fmt.Errorf("invalid user context")
	}

	userIDOpt := request.UserIDOption[string, bool](*ctxu.ID)

	req := request.New(
		ctx,
		register.Email,
		userIDOpt,
	)

	varifyEmail := models.AuthMsg{VerifyEmail: &req}

	if err := a.sendMsg(ctx, varifyEmail); err != nil {
		return nil, err
	}

	response, err := req.Wait()
	if err != nil {
		return nil, err
	}
	if response != nil && !*response {
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

	createUserIDOpt := request.UserIDOption[models.AuthUser, request.Signal](*ctxu.ID)

	creationRequest := request.New(ctx, creationUser, createUserIDOpt)

	createMsg := models.AuthMsg{
		RegisterUser: &creationRequest,
	}

	if err := a.sendMsg(ctx, createMsg); err != nil {
		return nil, err
	}

	_, err = creationRequest.Wait()
	if err != nil {
		return nil, err
	}

	return &models.RegisterResponse{
		Email:     creationUser.Email,
		FirstName: creationUser.FirstName,
		LastName:  creationUser.LastName,
	}, nil
}
