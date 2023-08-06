package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/repositories/database/sqlbuilder"
	"github.com/followthepattern/adapticc/pkg/request"
	"github.com/followthepattern/adapticc/pkg/utils/pointers"

	. "github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
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
	if requestBody.ID != nil {
		product, err := service.GetByID(userID, *requestBody.ID)
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
	product, err := service.Get(userID, requestBody)
	if err != nil {
		request.ReplyError(err)
		return
	}
	request.Reply(*product)
	return
}

func (service Product) replyCreate(req request.RequestHandler[[]models.Product, request.Signal]) {
	requestBody := req.RequestBody()
	userID := req.UserID()
	err := service.Create(userID, requestBody)
	if err != nil {
		req.ReplyError(err)
		return
	}
	req.Reply(request.Success())
	return
}

func (service Product) replyUpdate(req request.RequestHandler[models.Product, request.Signal]) {
	requestBody := req.RequestBody()
	userID := req.UserID()
	err := service.Update(userID, requestBody)
	if err != nil {
		req.ReplyError(err)
		return
	}
	req.Reply(request.Success())
	return
}

func (service Product) replyDelete(req request.RequestHandler[string, request.Signal]) {
	requestBody := req.RequestBody()
	userID := req.UserID()
	err := service.Delete(userID, requestBody)
	if err != nil {
		req.ReplyError(err)
		return
	}
	req.Reply(request.Success())
	return
}

func (repo Product) Create(userID string, products []models.Product) (err error) {
	count, err := sqlbuilder.GetInsertWithPermissions(repo.db, "PRODUCT", userID)
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

func (repo Product) GetByID(userID string, id string) (*models.Product, error) {
	product := models.Product{}

	query := repo.db.From(repo.tableName()).
		Where(Ex{
			"id": id})

	query = sqlbuilder.GetSelectWithPermissions(
		query,
		"PRODUCT",
		I("products.id"),
		userID,
	)

	_, err := query.ScanStruct(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo Product) Get(userID string, request models.ProductListRequestBody) (*models.ProductListResponse, error) {
	data := []models.Product{}

	query := repo.db.From(repo.tableName())

	if request.Filter.Search != nil {
		pattern := fmt.Sprintf("%%%s%%", *request.Filter.Search)
		query = query.Where(
			Or(
				I("id").Like(pattern),
				I("title").Like(pattern),
			))
	}

	query = sqlbuilder.GetSelectWithPermissions(
		query,
		"PRODUCT",
		I("products.id"),
		userID,
	)

	var count int64

	_, err := sqlbuilder.DistinctCount(query, I("products.id")).ScanVal(&count)
	if err != nil {
		return nil, err
	}

	if request.Pagination.Page == nil {
		request.Pagination.Page = pointers.ToPtr[uint](models.DefaultPage)
	}

	if request.Pagination.PageSize != nil {
		page := *request.Pagination.Page
		if page > 0 {
			page--
		}

		query = query.Offset(page * *request.Pagination.PageSize)
		query = query.Limit(*request.Pagination.PageSize)
	}

	orderLength := len(request.OrderBy)
	if orderLength > 0 {
		orderExpressions := make([]exp.OrderedExpression, orderLength)
		for i, order := range request.OrderBy {
			orderExpressions[i] = I(order.Name).Asc()
			if order.Desc != nil && *order.Desc {
				orderExpressions[i] = I(order.Name).Desc()
			}
		}
		query = query.Order(orderExpressions...)
	}

	err = query.Distinct().ScanStructs(&data)
	if err != nil {
		return nil, err
	}

	result := models.ProductListResponse{
		Count:    count,
		PageSize: request.Pagination.PageSize,
		Page:     request.Pagination.Page,
		Data:     data,
	}

	return &result, nil
}

func (repo Product) Update(userID string, request models.Product) error {
	query := repo.db.Update(repo.tableName()).
		Set(request)

	query = sqlbuilder.GetUpdateWithPermissions(
		query,
		"PRODUCT",
		I("products.id"),
		userID,
	)

	query = query.Where(I("id").Eq(*request.ID))

	_, err := query.
		Executor().
		Exec()
	return err
}

func (repo Product) Delete(userID, id string) error {
	query := repo.db.Delete(repo.tableName())

	query = sqlbuilder.GetDeleteWithPermissions(
		query,
		"PRODUCT",
		I("usr.products.id"),
		userID,
	)

	query = query.Where(C("id").Eq(id))

	_, err := query.
		Executor().
		Exec()

	return err
}
