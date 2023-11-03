package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/utils/pointers"

	. "github.com/followthepattern/goqu/v9"
	"github.com/followthepattern/goqu/v9/exp"
)

var productTable = S("usr").Table("products")

type Product struct {
	db *Database
}

func NewProduct(database *sql.DB) Product {
	db := New("postgres", database)

	return Product{
		db: db,
	}
}

func (repo Product) Create(products []models.Product) (err error) {
	for i, _ := range products {
		products[i].Userlog.CreatedAt = time.Now()
	}

	insertion := repo.db.Insert(productTable)

	_, err = insertion.Rows(products).Executor().Exec()
	return
}

func (repo Product) GetByID(id string) (*models.Product, error) {
	product := models.Product{}

	query := repo.db.From(productTable).
		Where(Ex{
			"id": id})

	_, err := query.ScanStruct(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo Product) Get(request models.ProductListRequestParams) (*models.ProductListResponse, error) {
	data := []models.Product{}

	query := repo.db.From(productTable)

	if len(request.Filter.Search) > 0 {
		pattern := fmt.Sprintf("%%%s%%", request.Filter.Search)
		query = query.Where(
			Or(
				I("id").Like(pattern),
				I("title").Like(pattern),
			))
	}

	count, err := query.Count()
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

	err = query.ScanStructs(&data)
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

func (repo Product) Update(model models.Product) error {
	model.Userlog.UpdatedAt = time.Now()

	_, err := repo.db.Update(productTable).
		Set(model).
		Where(I("id").Eq(model.ID)).
		Executor().
		Exec()

	return err
}

func (repo Product) Delete(id string) error {
	res, err := repo.db.
		Delete(productTable).
		Where(C("id").Eq(id)).
		Executor().
		Exec()

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows < 1 {
		return errors.New("no rows been deleted")
	}

	return err
}
