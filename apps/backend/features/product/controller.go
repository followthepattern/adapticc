package product

import (
	"context"
	"log/slog"

	"github.com/followthepattern/adapticc/config"
	"github.com/followthepattern/adapticc/container"
	"github.com/followthepattern/adapticc/features/auth"
	"github.com/followthepattern/adapticc/types"

	validation "github.com/go-ozzo/ozzo-validation"
)

type ProductController struct {
	product ProductService
	logger  *slog.Logger
	cfg     config.Config
}

func NewProductController(cont container.Container) ProductController {
	authorizationService := auth.NewAuthorizationService(cont, "product")
	productService := NewProductService(cont, authorizationService)

	return ProductController{
		product: productService,
		logger:  cont.GetLogger(),
		cfg:     cont.GetConfig(),
	}
}

func (ctrl ProductController) GetByID(ctx context.Context, id string) (*ProductModel, error) {
	if err := validation.Validate(id, types.Required("productID")); err != nil {
		return nil, err
	}

	result, err := ctrl.product.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if result.IsDefault() {
		return nil, nil
	}

	return result, nil
}

func (ctrl ProductController) Get(ctx context.Context, filter ProductListRequestParams) (*ProductListResponse, error) {
	return ctrl.product.Get(ctx, filter)
}

func (ctrl ProductController) Create(ctx context.Context, value ProductModel) error {
	if err := value.CreateValidate(); err != nil {
		return err
	}
	return ctrl.product.Create(ctx, value)
}

func (ctrl ProductController) Update(ctx context.Context, value ProductModel) error {
	if err := value.UpdateValidate(); err != nil {
		return err
	}

	return ctrl.product.Update(ctx, value)
}

func (ctrl ProductController) Delete(ctx context.Context, id string) error {
	if err := validation.Validate(id, types.Required("productID")); err != nil {
		return err
	}

	return ctrl.product.Delete(ctx, id)
}
