package product

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/followthepattern/adapticc/repositories/database/sqlbuilder"
	"github.com/followthepattern/adapticc/types"

	. "github.com/followthepattern/goqu/v9"
)

var ProductTable = S("usr").Table("products")

type ProductDatabase struct {
	db *Database
}

func NewProductDatabase(database *sql.DB) ProductDatabase {
	db := New("postgres", database)

	return ProductDatabase{
		db: db,
	}
}

func (repo ProductDatabase) Create(products []ProductModel) (err error) {
	for i, _ := range products {
		products[i].Userlog.CreatedAt = types.TimeNow()
	}

	insertion := repo.db.Insert(ProductTable)

	_, err = insertion.Rows(products).Executor().Exec()
	return
}

func (repo ProductDatabase) GetByID(id string) (*ProductModel, error) {
	product := ProductModel{}

	query := repo.db.From(ProductTable).
		Where(Ex{
			"id": id})

	_, err := query.ScanStruct(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo ProductDatabase) Get(request ProductListRequestParams) (*ProductListResponse, error) {
	data := []ProductModel{}

	query := repo.db.From(ProductTable)

	if request.Filter.Search.IsValid() {
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

	query = sqlbuilder.WithPagination(query, request.Pagination)

	query = sqlbuilder.WithOrderBy(query, request.OrderBy)

	err = query.ScanStructs(&data)
	if err != nil {
		return nil, err
	}

	result := ProductListResponse{
		Count:    types.Int64From(count),
		PageSize: request.Pagination.PageSize,
		Page:     request.Pagination.Page,
		Data:     data,
	}

	return &result, nil
}

func (repo ProductDatabase) Update(model ProductModel) error {
	model.Userlog.UpdatedAt = types.TimeNow()

	_, err := repo.db.Update(ProductTable).
		Set(model).
		Where(I("id").Eq(model.ID)).
		Executor().
		Exec()

	return err
}

func (repo ProductDatabase) Delete(id string) error {
	res, err := repo.db.
		Delete(ProductTable).
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
