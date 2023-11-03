package sqlexpectations

import (
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/followthepattern/adapticc/pkg/models"
)

func ExpectProduct(mock sqlmock.Sqlmock, result models.Product) {
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

	mock.ExpectQuery(sqlQuery).
		WillReturnRows(ModelToSQLMockRows(result))
}

func ExpectProducts(mock sqlmock.Sqlmock, filter models.ListFilter, page int, pageSize int, result []models.Product) {
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
			AddRow(len(result)))

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
		page*pageSize,
		pageSize,
	)

	mock.ExpectQuery(sqlQuery).
		WillReturnRows(ModelToSQLMockRows(result))
}

func CreateProduct(mock sqlmock.Sqlmock, userID string, product models.Product) {
	sqlQuery := fmt.Sprintf(`
		INSERT INTO "usr"."products"
			("created_at",
			"creation_user_id",
			"description",
			"id",
			"title")
		VALUES ('.*', '%s', '%s', '.*', '%s')`,
		userID,
		product.Description,
		product.Title,
	)

	mock.ExpectExec(sqlQuery).
		WillReturnResult(sqlmock.NewResult(1, 1))
}

func UpdateProduct(mock sqlmock.Sqlmock, userID string, product models.Product) {
	sqlQuery := fmt.Sprintf(`
	UPDATE
		"usr"."products"
	SET
		"description"='%s',"id"='%s',"title"='%s',"update_user_id"='%s',"updated_at"='.*' WHERE ("id" = '%s')`,
		product.Description,
		product.ID,
		product.Title,
		userID,
		product.ID)

	mock.ExpectExec(sqlQuery).
		WillReturnResult(sqlmock.NewResult(1, 1))
}

func DeleteProduct(mock sqlmock.Sqlmock, product models.Product) {
	sqlQuery := fmt.Sprintf(`DELETE FROM "usr"."products" WHERE ("id" = '%s')`,
		product.ID)

	mock.ExpectExec(sqlQuery).
		WillReturnResult(sqlmock.NewResult(1, 1))
}
