package controllers

import (
	"context"
	"fmt"

	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/request"
	"github.com/followthepattern/adapticc/pkg/services"
	"github.com/followthepattern/adapticc/pkg/utils"
	"github.com/followthepattern/adapticc/pkg/utils/pointers"
	"github.com/google/uuid"
)

type User struct {
	ctx               context.Context
	userMsgChannelOut chan<- models.UserMsg
	sendMsg           func(context.Context, models.UserMsg) error
}

func UserDependencyConstructor(cont *container.Container) (*User, error) {
	userMsgChannel, err := container.Resolve[services.UserMsgChannel](cont)
	if err != nil {
		return nil, err
	}

	return &User{
		ctx:               cont.GetContext(),
		userMsgChannelOut: *userMsgChannel,
		sendMsg:           request.CreateSenderFunc(*userMsgChannel, request.DefaultTimeOut),
	}, nil
}

func (ctrl User) GetByID(ctx context.Context, id string) (*models.User, error) {
	ctxu := utils.GetModelFromContext[models.User](ctx, utils.CtxUserKey)
	if ctxu == nil {
		return nil, fmt.Errorf("invalid user context")
	}

	userIDOpt := request.UserIDOption[models.UserRequestBody, models.User](*ctxu.ID)

	requestBody := models.UserRequestBody{ID: &id}

	req := request.New(
		ctx,
		requestBody,
		userIDOpt,
	)

	msg := models.UserMsg{Single: &req}

	if err := ctrl.sendMsg(ctx, msg); err != nil {
		return nil, err
	}

	result, err := req.Wait()
	if err != nil {
		return nil, err
	}

	if result.IsNil() {
		return nil, nil
	}

	return result, nil
}

func (ctrl User) Profile(ctx context.Context) (*models.User, error) {
	ctxu := utils.GetModelFromContext[models.User](ctx, utils.CtxUserKey)
	if ctxu.IsDefault() {
		return nil, fmt.Errorf("invalid user context")
	}

	requestBody := models.UserRequestBody{
		ID: ctxu.ID,
	}

	userIDOpt := request.UserIDOption[models.UserRequestBody, models.User](*ctxu.ID)

	req := request.New(
		ctx,
		requestBody,
		userIDOpt,
	)

	msg := models.UserMsg{Single: &req}

	if err := ctrl.sendMsg(ctx, msg); err != nil {
		return nil, err
	}

	user, err := req.Wait()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ctrl User) Get(ctx context.Context, filter models.UserListRequestBody) (*models.UserListResponse, error) {
	ctxu := utils.GetModelFromContext[models.User](ctx, utils.CtxUserKey)
	if ctxu == nil {
		return nil, fmt.Errorf("invalid user context")
	}

	userIDOpt := request.UserIDOption[models.UserListRequestBody, models.UserListResponse](*ctxu.ID)

	req := request.New(
		ctx,
		filter,
		userIDOpt,
	)

	msg := models.UserMsg{List: &req}

	if err := ctrl.sendMsg(ctx, msg); err != nil {
		return nil, err
	}

	result, err := req.Wait()
	if err != nil {
		return nil, err
	}

	response := models.UserListResponse(*result)

	return &response, nil
}

func (ctrl User) Create(ctx context.Context, user models.User) error {
	ctxu := utils.GetModelFromContext[models.User](ctx, utils.CtxUserKey)
	if ctxu == nil {
		return fmt.Errorf("invalid user context")
	}

	userIDOpt := request.UserIDOption[[]models.User, request.Signal](*ctxu.ID)

	user.ID = pointers.ToPtr(uuid.New().String())

	req := request.New(
		ctx,
		[]models.User{user},
		userIDOpt,
	)

	msg := models.UserMsg{
		Create: &req,
	}

	if err := ctrl.sendMsg(ctx, msg); err != nil {
		return err
	}

	_, err := req.Wait()

	return err
}

func (ctrl User) Update(ctx context.Context, user models.User) error {
	ctxu := utils.GetModelFromContext[models.User](ctx, utils.CtxUserKey)
	if ctxu == nil {
		return fmt.Errorf("invalid user context")
	}

	userIDOpt := request.UserIDOption[models.User, request.Signal](*ctxu.ID)

	req := request.New(
		ctx,
		user,
		userIDOpt,
	)

	msg := models.UserMsg{
		Update: &req,
	}

	if err := ctrl.sendMsg(ctx, msg); err != nil {
		return err
	}

	_, err := req.Wait()

	return err
}

func (ctrl User) Delete(ctx context.Context, id string) error {
	ctxu := utils.GetModelFromContext[models.User](ctx, utils.CtxUserKey)
	if ctxu == nil {
		return fmt.Errorf("invalid user context")
	}

	userIDOpt := request.UserIDOption[string, request.Signal](*ctxu.ID)

	req := request.New(
		ctx,
		id,
		userIDOpt,
	)

	msg := models.UserMsg{Delete: &req}

	if err := ctrl.sendMsg(ctx, msg); err != nil {
		return err
	}

	_, err := req.Wait()
	return err
}
