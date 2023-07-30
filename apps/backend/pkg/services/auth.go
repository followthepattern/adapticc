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
	userMsgChannelOut chan<- models.UserMsg
	ctx               context.Context
	cfg               config.Config
	sendMsg           func(ctx context.Context, msg models.UserMsg) error
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

	userMsgChannelOut, err := container.Resolve[repositories.UserMsgChannel](cont)
	if err != nil {
		return nil, err
	}

	dependency := Auth{
		ctx:               cont.GetContext(),
		cfg:               cont.GetConfig(),
		userMsgChannelOut: *userMsgChannelOut,
		authMsgChannelIn:  *authMsgChannelIn,
		sendMsg:           request.CreateSenderFunc(*userMsgChannelOut, request.DefaultTimeOut),
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
				result, err := service.Login(request.Login.Context(), request.Login.RequestBody().Email, request.Login.RequestBody().Password)
				if err != nil {
					request.Login.ReplyError(err)
					continue
				}
				request.Login.Reply(*result)
			} else if request.Register != nil {
				register := request.Register.RequestBody()
				result, err := service.Register(request.Register.Context(), register.Email, register.FirstName, register.LastName, register.Password)
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
	requestBody := models.UserRequestBody{
		Email: pointers.String(email),
	}

	req := request.New[models.UserRequestBody, models.User](ctx, requestBody)

	msg := models.UserMsg{Single: &req}

	if err := a.sendMsg(ctx, msg); err != nil {
		return nil, err
	}

	user, err := req.Wait()
	if err != nil {
		return nil, err
	}

	if user == nil || user.ID == nil {
		return nil, errors.New(WRONG_EMAIL_OR_PASSWORD)
	}

	requestPasswordHash := utils.GeneratePasswordHash(password, *user.Salt)

	if requestPasswordHash != *user.Password {
		return nil, errors.New(WRONG_EMAIL_OR_PASSWORD)
	}

	expiresAt := time.Now().Add(time.Hour * 24)
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"ID":        *user.ID,
		"email":     *user.Email,
		"firstName": *user.FirstName,
		"lastName":  *user.LastName,
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

func (a Auth) Register(ctx context.Context, email string, firstName string, lastName string, password string) (*models.RegisterResponse, error) {
	requestBody := models.UserRequestBody{Email: pointers.String(email)}

	ctxu := utils.GetModelFromContext[models.User](ctx, utils.CtxUserKey)
	if ctxu == nil {
		return nil, fmt.Errorf("invalid user context")
	}

	userIDOpt := request.UserIDOption[models.UserRequestBody, models.User](*ctxu.ID)

	req := request.New(
		ctx,
		requestBody,
		userIDOpt,
	)

	singleMsg := models.UserMsg{Single: &req}

	if err := a.sendMsg(ctx, singleMsg); err != nil {
		return nil, err
	}

	response, err := req.Wait()
	if err != nil {
		return nil, err
	}
	if response != nil && !response.IsNil() {
		return nil, fmt.Errorf(EMAIL_IS_ALREADY_IN_USE_PATTERN, email)
	}

	salt := utils.GenerateSaltString()
	passwordHash := utils.GeneratePasswordHash(password, salt)

	creationUser := models.User{
		ID:        pointers.String(uuid.New().String()),
		Email:     &email,
		FirstName: &firstName,
		LastName:  &lastName,
		Password:  &passwordHash,
		Salt:      &salt,
		Active:    pointers.Bool(true),
	}

	createUserIDOpt := request.UserIDOption[[]models.User, request.Signal](*ctxu.ID)

	creationRequest := request.New(ctx, []models.User{creationUser}, createUserIDOpt)

	createMsg := models.UserMsg{
		Create: &creationRequest,
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
