package rest

import (
	"encoding/json"
	"net/http"

	"github.com/followthepattern/adapticc/pkg/controllers"
	"github.com/followthepattern/adapticc/pkg/models"
)

type Product struct {
	product controllers.Product
}

func NewProduct(ctrl controllers.Product) Product {
	return Product{
		product: ctrl,
	}
}

func (service Product) Create(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		BadRequest(w, FailedToDecodeRequestBody)
		return
	}

	err = service.product.Create(r.Context(), product)
	if err != nil {
		BadRequest(w, err.Error())
		return
	}

	Success(w, "Created")
}
