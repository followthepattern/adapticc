package product

import (
	"encoding/json"
	"net/http"

	"github.com/followthepattern/adapticc/api"
)

type ProductRest struct {
	product ProductController
}

func NewProductRest(ctrl ProductController) ProductRest {
	return ProductRest{
		product: ctrl,
	}
}

func (service ProductRest) Create(w http.ResponseWriter, r *http.Request) {
	var product ProductModel
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		api.BadRequest(w, api.FailedToDecodeRequestBody)
		return
	}

	err = service.product.Create(r.Context(), product)
	if err != nil {
		api.BadRequest(w, err.Error())
		return
	}

	api.Created(w)
}
