package product

import (
	"encoding/json"
	"net/http"

	"github.com/followthepattern/adapticc/api/httpresponses"
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
		httpresponses.BadRequest(w, httpresponses.FailedToDecodeRequestBody)
		return
	}

	err = service.product.Create(r.Context(), product)
	if err != nil {
		httpresponses.BadRequest(w, err.Error())
		return
	}

	httpresponses.Created(w)
}
