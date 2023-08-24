package rest

import (
	"encoding/json"
	"net/http"

	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/followthepattern/adapticc/pkg/controllers"
	"github.com/followthepattern/adapticc/pkg/models"
	"go.uber.org/zap"
)

type Product struct {
	product *controllers.Product
	logger  *zap.Logger
}

func newProduct(cont *container.Container) (*Product, error) {
	service, err := container.Resolve[controllers.Product](cont)
	if err != nil {
		return nil, err
	}

	return &Product{
		product: service,
		logger:  cont.GetLogger(),
	}, nil
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
