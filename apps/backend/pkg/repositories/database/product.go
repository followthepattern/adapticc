package database

import (
	"context"
	"errors"

	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/repositories/database/sqlbuilder"
	"github.com/followthepattern/adapticc/pkg/request"
	"github.com/followthepattern/adapticc/pkg/utils/pointers"

	. "github.com/doug-martin/goqu/v9"
)

type ProductMsgChannel chan models.ProductMsg

func RegisterProductChannel(cont *container.Container) {
	if cont == nil {
		return
	}
	requestChan := make(ProductMsgChannel)
	container.Register(cont, func(cont *container.Container) (*ProductMsgChannel, error) {
		return &requestChan, nil
	})
}

type Product struct {
	usrMsgIn <-chan models.ProductMsg
	db       Database
	ctx      context.Context
}

func (Product) tableName() string {
	return "usr.products"
}

func ProductDependencyConstructor(cont *container.Container) (*Product, error) {
	db := New("postgres", cont.GetDB())

	if db == nil {
		return nil, errors.New("db is null")
	}

	userMsg, err := container.Resolve[ProductMsgChannel](cont)
	if err != nil {
		return nil, err
	}

	dependency := &Product{
		ctx:      cont.GetContext(),
		db:       *db,
		usrMsgIn: *userMsg,
	}

	go func() {
		dependency.MonitorChannels()
	}()

	return dependency, nil
}

func (repository Product) MonitorChannels() {
	for {
		select {
		case request := <-repository.usrMsgIn:
			repository.replyRequest(request)
		case <-repository.ctx.Done():
			return
		}
	}
}

func (service Product) replyRequest(request models.ProductMsg) {
	switch {
	case request.Single != nil:
		service.replySingle(*request.Single)
	case request.List != nil:
		service.replyList(*request.List)
	case request.Create != nil:
		service.replyCreate(*request.Create)
	case request.Update != nil:
		service.replyUpdate(*request.Update)
	case request.Delete != nil:
		service.replyDelete(*request.Delete)
	}
}

func (service Product) replySingle(request request.RequestHandler[models.ProductRequestBody, models.Product]) {
	requestBody := request.RequestBody()
	userID := request.UserID()
	if requestBody.ProductID != nil && userID != nil {
		product, err := service.GetByID(*userID, *requestBody.ProductID)
		if err != nil {
			request.ReplyError(err)
			return
		}
		request.Reply(*product)
		return
	}

	request.ReplyError(errors.New("missing userID or productID"))
}

func (service Product) replyList(request request.RequestHandler[models.ProductListRequestBody, models.ProductListResponse]) {
	requestBody := request.RequestBody()
	userID := request.UserID()
	if userID != nil {
		product, err := service.Get(*userID, requestBody)
		if err != nil {
			request.ReplyError(err)
			return
		}
		request.Reply(*product)
		return
	}

	request.ReplyError(errors.New("missing userID"))
}

func (service Product) replyCreate(req request.RequestHandler[[]models.Product, request.Signal]) {
	requestBody := req.RequestBody()
	userID := req.UserID()
	if userID != nil {
		err := service.Create(*userID, requestBody)
		if err != nil {
			req.ReplyError(err)
			return
		}
		req.Reply(request.Success())
		return
	}

	req.ReplyError(errors.New("missing userID"))
}

func (service Product) replyUpdate(req request.RequestHandler[models.Product, request.Signal]) {
	requestBody := req.RequestBody()
	userID := req.UserID()
	if userID != nil {
		err := service.Update(*userID, requestBody)
		if err != nil {
			req.ReplyError(err)
			return
		}
		req.Reply(request.Success())
		return
	}

	req.ReplyError(errors.New("missing userID or productID"))
}

func (service Product) replyDelete(req request.RequestHandler[string, request.Signal]) {
	requestBody := req.RequestBody()
	userID := req.UserID()
	if userID != nil {
		err := service.Delete(*userID, requestBody)
		if err != nil {
			req.ReplyError(err)
			return
		}
		req.Reply(request.Success())
		return
	}

	req.ReplyError(errors.New("missing userID or productID"))
}

func (repo Product) Create(userID string, products []models.Product) (err error) {
	count, err := sqlbuilder.GetInsertWithPermissions(repo.db, "product", userID)
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("there is no effective permission to create this resource")
	}

	// TODO: set log fields
	// for i, _ := range products {
	// 	products[i].CreatedAt = pointers.Time(time.Now())
	// }

	insertion := repo.db.Insert(repo.tableName())

	_, err = insertion.Rows(products).Executor().Exec()
	return
}

func (repo Product) GetByID(userID string, productID string) (*models.Product, error) {
	product := models.Product{}

	query := repo.db.From(repo.tableName()).
		Where(Ex{
			"product_id": productID})

	query = sqlbuilder.GetSelectWithPermissions(
		query,
		"product",
		I("products.product_id"),
		userID,
	)

	_, err := query.ScanStruct(&product)
	if err != nil {
		return nil, err
	}
	return &product, err
}

func (repo Product) Get(userID string, request models.ProductListRequestBody) (*models.ProductListResponse, error) {
	entities := []models.Product{}

	query := repo.db.From(repo.tableName())

	query = sqlbuilder.GetSelectWithPermissions(
		query,
		"product",
		I("products.product_id"),
		userID,
	)

	var count int64

	_, err := sqlbuilder.DistinctCount(query, I("products.product_id")).ScanVal(&count)
	if err != nil {
		return nil, err
	}

	result := models.ProductListResponse{
		Count:    count,
		PageSize: request.PageSize,
	}

	if request.Page == nil {
		result.Page = pointers.UInt(1)
	}

	if request.PageSize != nil {
		page := *request.Page
		if page > 0 {
			page--
		}

		query = query.Offset(page * *request.PageSize)
		query = query.Limit(*request.PageSize)
	}

	err = query.Distinct().ScanStructs(&entities)
	if err != nil {
		return nil, err
	}

	result.Data = entities

	return &result, nil
}

func (repo Product) Update(userID string, request models.Product) error {
	query := repo.db.Update(repo.tableName()).
		Set(request)

	query = sqlbuilder.GetUpdateWithPermissions(
		query,
		"product",
		I("products.product_id"),
		userID,
	)

	query = query.Where(I("product_id").Eq(*request.ProductID))

	_, err := query.
		Executor().
		Exec()
	return err
}

func (repo Product) Delete(userID, productID string) error {
	query := repo.db.Delete(repo.tableName())

	query = sqlbuilder.GetDeleteWithPermissions(
		query,
		"product",
		I("products.product_id"),
		userID,
	)

	query = query.Where(C("product_id").Eq(productID))

	_, err := query.
		Executor().
		Exec()

	return err
}
