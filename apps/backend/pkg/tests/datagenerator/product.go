package datagenerator

import (
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/utils/pointers"
)

func NewRandomProduct() models.Product {
	return models.Product{
		ProductID:   pointers.String(String(8)),
		Title:       pointers.String(String(8)),
		Description: pointers.String(String(8)),
	}
}
