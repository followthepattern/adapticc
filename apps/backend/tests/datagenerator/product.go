package datagenerator

import (
	"github.com/followthepattern/adapticc/features/product"
	"github.com/followthepattern/adapticc/types"
)

func NewRandomProduct() product.ProductModel {
	return product.ProductModel{
		ID:          types.StringFrom(String(8)),
		Title:       types.StringFrom(String(8)),
		Description: types.StringFrom(String(8)),
	}
}
