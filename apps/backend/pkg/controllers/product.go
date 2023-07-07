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

type Product struct {
	productMsgChannelOut chan<- models.ProductMsg
	sendMsg              func(ctx context.Context, msg models.ProductMsg) error
}

func ProductDependencyConstructor(cont *container.Container) (*Product, error) {
	productMsgChannelOut, err := container.Resolve[services.ProductMsgChannel](cont)
	if err != nil {
		return nil, err
	}

	dependency := Product{
		productMsgChannelOut: *productMsgChannelOut,
		sendMsg:              request.CreateSenderFunc(*productMsgChannelOut, request.DefaultTimeOut),
	}

	return &dependency, nil
}

func (ctrl Product) GetByID(ctx context.Context, id string) (*models.Product, error) {
	ctxu := utils.GetModelFromContext[models.User](ctx, utils.CtxUserKey)
	if ctxu == nil {
		return nil, fmt.Errorf("invalid user context")
	}

	userIDOpt := request.UserIDOption[models.ProductRequestBody, models.Product](*ctxu.ID)

	requestBody := models.ProductRequestBody{ID: &id}

	req := request.New(
		ctx,
		requestBody,
		userIDOpt,
	)

	msg := models.ProductMsg{Single: &req}

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

func (ctrl Product) Get(ctx context.Context, filter models.ProductListRequestBody) (*models.ProductListResponse, error) {
	ctxu := utils.GetModelFromContext[models.User](ctx, utils.CtxUserKey)
	if ctxu == nil {
		return nil, fmt.Errorf("invalid user context")
	}

	userIDOpt := request.UserIDOption[models.ProductListRequestBody, models.ProductListResponse](*ctxu.ID)

	req := request.New(
		ctx,
		filter,
		userIDOpt,
	)

	msg := models.ProductMsg{List: &req}

	if err := ctrl.sendMsg(ctx, msg); err != nil {
		return nil, err
	}

	result, err := req.Wait()
	if err != nil {
		return nil, err
	}

	response := models.ProductListResponse(*result)

	return &response, nil
}

func (ctrl Product) Create(ctx context.Context, value models.Product) error {
	ctxu := utils.GetModelFromContext[models.User](ctx, utils.CtxUserKey)
	if ctxu == nil {
		return fmt.Errorf("invalid user context")
	}

	userIDOpt := request.UserIDOption[[]models.Product, request.Signal](*ctxu.ID)

	value.ID = pointers.ToPtr(uuid.New().String())

	req := request.New(
		ctx,
		[]models.Product{value},
		userIDOpt,
	)

	msg := models.ProductMsg{
		Create: &req,
	}

	if err := ctrl.sendMsg(ctx, msg); err != nil {
		return err
	}

	_, err := req.Wait()

	return err
}

func (ctrl Product) Update(ctx context.Context, value models.Product) error {
	ctxu := utils.GetModelFromContext[models.User](ctx, utils.CtxUserKey)
	if ctxu == nil {
		return fmt.Errorf("invalid user context")
	}

	userIDOpt := request.UserIDOption[models.Product, request.Signal](*ctxu.ID)

	req := request.New(
		ctx,
		value,
		userIDOpt,
	)

	msg := models.ProductMsg{
		Update: &req,
	}

	if err := ctrl.sendMsg(ctx, msg); err != nil {
		return err
	}

	_, err := req.Wait()

	return err
}

func (ctrl Product) Delete(ctx context.Context, id string) error {
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

	msg := models.ProductMsg{Delete: &req}

	if err := ctrl.sendMsg(ctx, msg); err != nil {
		return err
	}

	_, err := req.Wait()
	return err
}
