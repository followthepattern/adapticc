package datagenerator

import (
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/utils/pointers"
)

func NewRandomProduct() models.Product {
	return models.Product{
		ID:          pointers.ToPtr(String(8)),
		Title:       pointers.ToPtr(String(8)),
		Description: pointers.ToPtr(String(8)),
	}
}