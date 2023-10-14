package controllers

import (
	"context"
	"database/sql"

	"github.com/followthepattern/adapticc/pkg/accesscontrol"
	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/repositories/database"
	"github.com/followthepattern/adapticc/pkg/services"
	"go.uber.org/zap"
)

type Product struct {
	product services.Product
	logger  *zap.Logger
	cfg     config.Config
}

func NewProduct(ctx context.Context, ac accesscontrol.AccessControl, db *sql.DB, cfg config.Config, logger *zap.Logger) Product {
	productRepository := database.NewProduct(ctx, db)
	roleRepository := database.NewRole(ctx, db)
	productService := services.NewProduct(ctx, ac, productRepository, roleRepository, cfg, logger)

	return Product{
		product: productService,
		logger:  logger,
		cfg:     cfg,
	}
}

func (ctrl Product) GetByID(ctx context.Context, id string) (*models.Product, error) {
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
	return ctrl.product.Delete(ctx, id)
}