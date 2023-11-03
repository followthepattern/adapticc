package sqlexpectations

import (
	"fmt"

	"github.com/followthepattern/adapticc/pkg/models"

	"github.com/DATA-DOG/go-sqlmock"
)

func ExpectGetUserByEmail(mock sqlmock.Sqlmock, result models.User, email string) {
	sqlQuery := fmt.Sprintf(`SELECT "active", "created_at", "creation_user_id", "email", "first_name", "id", "last_name", "update_user_id", "updated_at" FROM "usr"."users" WHERE ("email" = '%v') LIMIT 1`, email)

	mock.ExpectQuery(sqlQuery).
		WillReturnRows(ModelToSQLMockRows(result))
}

func ExpectGetUserByID(mock sqlmock.Sqlmock, result models.User, id string) {
	sqlQuery := fmt.Sprintf(`SELECT "active", "created_at", "creation_user_id", "email", "first_name", "id", "last_name", "update_user_id", "updated_at" FROM "usr"."users" WHERE ("id" = '%v') LIMIT 1`, id)

	mock.ExpectQuery(sqlQuery).
		WillReturnRows(ModelToSQLMockRows(result))
}

func ExpectUsers(mock sqlmock.Sqlmock, result []models.User, listRequestParams models.UserListRequestParams) {
	countQuery := fmt.Sprintf(`
	SELECT
		COUNT(\*) AS "count"
	FROM
		"usr"."users"
	WHERE
		(("first_name" LIKE '%%%s%%') OR
		("last_name" LIKE '%%%s%%') OR
		("email" LIKE '%%%s%%')) LIMIT 1`,
		*listRequestParams.Filter.Search,
		*listRequestParams.Filter.Search,
		*listRequestParams.Filter.Search)

	mock.ExpectQuery(countQuery).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(len(result)))

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
		*listRequestParams.Filter.Search,
		*listRequestParams.Filter.Search,
		*listRequestParams.Filter.Search,
		*listRequestParams.Pagination.PageSize)

	SQLMockRows := ModelToSQLMockRows(result)
	mock.ExpectQuery(sqlQuery).
		WillReturnRows(SQLMockRows)
}

func ExpectUsersWithoutPaging(mock sqlmock.Sqlmock, result []models.User, listRequestParams models.UserListRequestParams) {
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
		*listRequestParams.Filter.Search,
		*listRequestParams.Filter.Search,
		*listRequestParams.Filter.Search)

	mock.ExpectQuery(countQuery).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(len(result)))

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
		*listRequestParams.Filter.Search,
		*listRequestParams.Filter.Search,
		*listRequestParams.Filter.Search)

	SQLMockRows := ModelToSQLMockRows(result)
	mock.ExpectQuery(sqlQuery).
		WillReturnRows(SQLMockRows)
}

func ExpectCreateUser(mock sqlmock.Sqlmock, userID string, insert models.User) {
	sqlQuery := fmt.Sprintf(`
	INSERT INTO
		"usr"."users" ("created_at",
			"creation_user_id",
			"email",
			"first_name",
			"id",
			"last_name")
		VALUES ('.*', '%s', '%s', '%s', '.*', '%s')`,
		userID,
		insert.Email,
		insert.FirstName,
		insert.LastName,
	)

	mock.ExpectExec(sqlQuery).
		WillReturnResult(sqlmock.NewResult(1, 1))
}

func ExpectUpdateUser(mock sqlmock.Sqlmock, userID string, model models.User) {
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

func ExpectDeleteUser(mock sqlmock.Sqlmock, userID string) {
	sqlQuery := fmt.Sprintf(`
	DELETE FROM
		"usr"."users"
	WHERE
		("id" = '%s')`,
		userID)

	mock.ExpectExec(sqlQuery).
		WillReturnResult(sqlmock.NewResult(1, 1))
}
