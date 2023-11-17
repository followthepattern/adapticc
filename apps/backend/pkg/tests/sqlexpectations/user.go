package sqlexpectations

import (
	"database/sql/driver"
	"fmt"

	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/types"

	"github.com/DATA-DOG/go-sqlmock"
)

var userColumns = []string{
	"active",
	"created_at",
	"creation_user_id",
	"email",
	"first_name",
	"id",
	"last_name",
	"update_user_id",
	"updated_at"}

func ExpectGetUserByEmail(mock sqlmock.Sqlmock, result models.User, email types.String) {
	sqlQuery := fmt.Sprintf(`SELECT "active", "created_at", "creation_user_id", "email", "first_name", "id", "last_name", "update_user_id", "updated_at" FROM "usr"."users" WHERE ("email" = '%v') LIMIT 1`, email)

	mock.ExpectQuery(sqlQuery).
		WillReturnRows(ModelToSQLMockRows(result))
}

func ExpectGetUserByID(mock sqlmock.Sqlmock, result models.User, id types.String) {
	sqlQuery := fmt.Sprintf(`SELECT "active", "created_at", "creation_user_id", "email", "first_name", "id", "last_name", "update_user_id", "updated_at" FROM "usr"."users" WHERE ("id" = '%v') LIMIT 1`, id)

	rows := sqlmock.NewRows(userColumns)

	values := []driver.Value{
		result.Active,
		result.CreatedAt,
		result.CreationUserID,
		result.Email,
		result.FirstName,
		result.ID,
		result.LastName,
		result.UpdateUserID,
		result.UpdatedAt,
	}
	rows.AddRow(values...)

	mock.ExpectQuery(sqlQuery).
		WillReturnRows(rows)
}

func ExpectUsers(mock sqlmock.Sqlmock, results []models.User, listRequestParams models.UserListRequestParams) {
	countQuery := fmt.Sprintf(`
	SELECT
		COUNT(\*) AS "count"
	FROM
		"usr"."users"
	WHERE
		(("first_name" LIKE '%%%s%%') OR
		("last_name" LIKE '%%%s%%') OR
		("email" LIKE '%%%s%%')) LIMIT 1`,
		listRequestParams.Filter.Search,
		listRequestParams.Filter.Search,
		listRequestParams.Filter.Search)

	mock.ExpectQuery(countQuery).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(len(results)))

	sqlQuery := fmt.Sprintf(`
	SELECT
		"active",
		"created_at",
		"creation_user_id",
		"email",
		"first_name",
		"id",
		"last_name",
		"update_user_id",
		"updated_at"
	FROM
		"usr"."users"
	WHERE
		(("first_name" LIKE '%%%s%%') OR
		("last_name" LIKE '%%%s%%') OR
		("email" LIKE '%%%s%%')) LIMIT %v`,
		listRequestParams.Filter.Search,
		listRequestParams.Filter.Search,
		listRequestParams.Filter.Search,
		listRequestParams.Pagination.PageSize)

	rows := sqlmock.NewRows(userColumns)

	for _, result := range results {
		values := []driver.Value{
			result.Active,
			result.CreatedAt,
			result.CreationUserID,
			result.Email,
			result.FirstName,
			result.ID,
			result.LastName,
			result.UpdateUserID,
			result.UpdatedAt,
		}
		rows.AddRow(values...)
	}

	mock.ExpectQuery(sqlQuery).
		WillReturnRows(rows)
}

func ExpectUsersWithoutPaging(mock sqlmock.Sqlmock, results []models.User, listRequestParams models.UserListRequestParams) {
	countQuery := fmt.Sprintf(`
	SELECT
		COUNT(\*) AS "count"
	FROM
		"usr"."users"
	WHERE
		(("first_name" LIKE '%%%s%%') OR
		("last_name" LIKE '%%%s%%') OR
		("email" LIKE '%%%s%%'))
	LIMIT 1`,
		listRequestParams.Filter.Search,
		listRequestParams.Filter.Search,
		listRequestParams.Filter.Search)

	mock.ExpectQuery(countQuery).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(len(results)))

	sqlQuery := fmt.Sprintf(`
	SELECT
		"active",
		"created_at",
		"creation_user_id",
		"email",
		"first_name",
		"id",
		"last_name",
		"update_user_id",
		"updated_at"
	FROM
		"usr"."users"
	WHERE
		(("first_name" LIKE '%%%s%%') OR
		("last_name" LIKE '%%%s%%') OR
		("email" LIKE '%%%s%%'))`,
		listRequestParams.Filter.Search,
		listRequestParams.Filter.Search,
		listRequestParams.Filter.Search)

	rows := sqlmock.NewRows(userColumns)

	for _, result := range results {
		values := []driver.Value{
			result.Active,
			result.CreatedAt,
			result.CreationUserID,
			result.Email,
			result.FirstName,
			result.ID,
			result.LastName,
			result.UpdateUserID,
			result.UpdatedAt,
		}
		rows.AddRow(values...)
	}

	mock.ExpectQuery(sqlQuery).
		WillReturnRows(rows)
}

func ExpectCreateUser(mock sqlmock.Sqlmock, userID types.String, insert models.User) {
	sqlQuery := fmt.Sprintf(`
	INSERT INTO
		"usr"."users" ("active",
			"created_at",
			"creation_user_id",
			"email",
			"first_name",
			"id",
			"last_name")
		VALUES (FALSE, '.*', '%s', '%s', '%s', '.*', '%s')`,
		userID,
		insert.Email,
		insert.FirstName,
		insert.LastName,
	)

	mock.ExpectExec(sqlQuery).
		WillReturnResult(sqlmock.NewResult(1, 1))
}

func ExpectUpdateUser(mock sqlmock.Sqlmock, userID types.String, model models.User) {
	sqlQuery := fmt.Sprintf(`
	UPDATE
		"usr"."users"
	SET
		"first_name"='%s',"last_name"='%s',"update_user_id"='%s',"updated_at"='.*'
	WHERE
		("id" = '%s')`,
		model.FirstName,
		model.LastName,
		userID,
		model.ID,
	)

	mock.ExpectExec(sqlQuery).
		WillReturnResult(sqlmock.NewResult(1, 1))
}

func ExpectDeleteUser(mock sqlmock.Sqlmock, userID types.String) {
	sqlQuery := fmt.Sprintf(`
	DELETE FROM
		"usr"."users"
	WHERE
		("id" = '%s')`,
		userID)

	mock.ExpectExec(sqlQuery).
		WillReturnResult(sqlmock.NewResult(1, 1))
}
