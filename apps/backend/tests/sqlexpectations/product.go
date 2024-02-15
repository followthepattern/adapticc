package sqlexpectations

import (
	"database/sql/driver"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/followthepattern/adapticc/features/product"
	"github.com/followthepattern/adapticc/models"
	"github.com/followthepattern/adapticc/types"
)

var productColumns = []string{
	"created_at",
	"creation_user_id",
	"description",
	"id",
	"title",
	"update_user_id",
	"updated_at"}

func ExpectProduct(mock sqlmock.Sqlmock, result product.ProductModel) {
	sqlQuery := fmt.Sprintf(
		`SELECT
			"created_at",
			"creation_user_id",
			"description",
			"id",
			"title",
			"update_user_id",
			"updated_at"
		FROM
			"usr"."products"
		WHERE
			("id" = '%s')
		LIMIT 1`,
		result.ID)

	rows := sqlmock.NewRows(productColumns)
	values := []driver.Value{
		result.CreatedAt,
		result.CreationUserID,
		result.Description,
		result.ID,
		result.Title,
		result.UpdateUserID,
		result.UpdatedAt,
	}
	rows.AddRow(values...)

	mock.ExpectQuery(sqlQuery).
		WillReturnRows(rows)
}

func ExpectProducts(mock sqlmock.Sqlmock, filter models.ListFilter, page int, pageSize int, results []product.ProductModel) {
	countQuery := fmt.Sprintf(`
	SELECT
		COUNT(\*) AS "count"
	FROM
		"usr"."products"
	WHERE (("id" LIKE '%%%s%%') OR ("title" LIKE '%%%s%%'))
	LIMIT 1`,
		filter.Search,
		filter.Search,
	)

	mock.ExpectQuery(countQuery).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(len(results)))

	if page > 0 {
		page--
	}

	sqlQuery := fmt.Sprintf(`
	SELECT
		"created_at",
		"creation_user_id",
		"description",
		"id",
		"title",
		"update_user_id",
		"updated_at"
	FROM "usr"."products"
	WHERE (("id" LIKE '%%%s%%') OR ("title" LIKE '%%%s%%'))
	LIMIT %d OFFSET %d`,
		filter.Search,
		filter.Search,
		pageSize,
		page*pageSize,
	)

	rows := sqlmock.NewRows(productColumns)

	for _, result := range results {
		values := []driver.Value{
			result.CreatedAt,
			result.CreationUserID,
			result.Description,
			result.ID,
			result.Title,
			result.UpdateUserID,
			result.UpdatedAt,
		}
		rows.AddRow(values...)
	}

	mock.ExpectQuery(sqlQuery).
		WillReturnRows(rows)
}

func CreateProduct(mock sqlmock.Sqlmock, userID types.String, product product.ProductModel) {
	sqlQuery := fmt.Sprintf(`
		INSERT INTO "usr"."products"
			("created_at",
			"creation_user_id",
			"description",
			"id",
			"title")
		VALUES ('.*', '%s', '%s', '.*', '%s')`,
		userID,
		product.Description.Data,
		product.Title.Data,
	)

	mock.ExpectExec(sqlQuery).
		WillReturnResult(sqlmock.NewResult(1, 1))
}

func UpdateProduct(mock sqlmock.Sqlmock, userID types.String, product product.ProductModel) {
	sqlQuery := fmt.Sprintf(`
	UPDATE
		"usr"."products"
	SET
		"description"='%s',"id"='%s',"title"='%s',"update_user_id"='%s',"updated_at"='.*' WHERE ("id" = '%s')`,
		product.Description.Data,
		product.ID.Data,
		product.Title.Data,
		userID,
		product.ID.Data)

	mock.ExpectExec(sqlQuery).
		WillReturnResult(sqlmock.NewResult(1, 1))
}

func DeleteProduct(mock sqlmock.Sqlmock, product product.ProductModel) {
	sqlQuery := fmt.Sprintf(`DELETE FROM "usr"."products" WHERE ("id" = '%s')`,
		product.ID.Data)

	mock.ExpectExec(sqlQuery).
		WillReturnResult(sqlmock.NewResult(1, 1))
}
