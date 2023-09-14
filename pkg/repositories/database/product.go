package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/repositories/database/sqlbuilder"
	"github.com/followthepattern/adapticc/pkg/utils/pointers"

	. "github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

const productTable = "usr.products"

type Product struct {
	db  *Database
	ctx context.Context
}

func NewProduct(ctx context.Context, database *sql.DB) Product {
	db := New("postgres", database)

	return Product{
		ctx: ctx,
		db:  db,
	}
}

func (repo Product) Create(userID string, products []models.Product) (err error) {
	count, err := sqlbuilder.GetInsertWithPermissions(repo.db, "PRODUCT", userID)
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("there is no effective permission to create this resource")
	}

	for i, _ := range products {
		products[i].Userlog = setCreateUserlog(userID, time.Now())
	}

	insertion := repo.db.Insert(productTable)

	_, err = insertion.Rows(products).Executor().Exec()
	return
}

func (repo Product) GetByID(userID string, id string) (*models.Product, error) {
	product := models.Product{}

	query := repo.db.From(productTable).
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

func (repo Product) Get(userID string, request models.ProductListRequestParams) (*models.ProductListResponse, error) {
	data := []models.Product{}

	query := repo.db.From(productTable)

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

func (repo Product) Update(userID string, model models.Product) error {
	model.Userlog = setUpdateUserlog(userID, time.Now())

	query := repo.db.Update(productTable).
		Set(model)

	query = sqlbuilder.GetUpdateWithPermissions(
		query,
		"PRODUCT",
		I("products.id"),
		userID,
	)

	query = query.Where(I("id").Eq(*model.ID))

	_, err := query.
		Executor().
		Exec()
	return err
}

func (repo Product) Delete(userID, id string) error {
	query := repo.db.Delete(productTable)

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
