package controllers

import (
	"context"
	"log/slog"

	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/repositories/database"
	"github.com/followthepattern/adapticc/pkg/services"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Product struct {
	product services.Product
	logger  *slog.Logger
	cfg     config.Config
}

func NewProduct(cont container.Container) Product {
	productRepository := database.NewProduct(cont.GetDB())
	roleRepository := database.NewRole(cont.GetDB())
	productService := services.NewProduct(cont, productRepository, roleRepository)

	return Product{
		product: productService,
		logger:  cont.GetLogger(),
		cfg:     cont.GetConfig(),
	}
}

func (ctrl Product) GetByID(ctx context.Context, id string) (*models.Product, error) {
	if err := validation.Validate(id, Required("productID")); err != nil {
		return nil, err
	}

	result, err := ctrl.product.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if result.IsNil() {
		return nil, nil
	}

	return result, nil
}

func (ctrl Product) Get(ctx context.Context, filter models.ProductListRequestParams) (*models.ProductListResponse, error) {
	return ctrl.product.Get(ctx, filter)
}

func (ctrl Product) Create(ctx context.Context, value models.Product) error {
	if err := value.CreateValidate(); err != nil {
		return err
	}
	return ctrl.product.Create(ctx, value)
}

func (ctrl Product) Update(ctx context.Context, value models.Product) error {
	if err := value.UpdateValidate(); err != nil {
		return err
	}

	return ctrl.product.Update(ctx, value)
}

func (ctrl Product) Delete(ctx context.Context, id string) error {
	if err := validation.Validate(id, Required("productID")); err != nil {
		return err
	}

	return ctrl.product.Delete(ctx, id)
}
