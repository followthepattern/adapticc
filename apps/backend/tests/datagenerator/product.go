package datagenerator

import (
	"github.com/followthepattern/adapticc/models"
	"github.com/followthepattern/adapticc/types"
)

func NewRandomProduct() models.Product {
	return models.Product{
		ID:          types.StringFrom(String(8)),
		Title:       types.StringFrom(String(8)),
		Description: types.StringFrom(String(8)),
	}
}
