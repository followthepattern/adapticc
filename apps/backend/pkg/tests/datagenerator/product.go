package datagenerator

import (
	"github.com/followthepattern/adapticc/pkg/models"
)

func NewRandomProduct() models.Product {
	return models.Product{
		ID:          String(8),
		Title:       String(8),
		Description: String(8),
	}
}
