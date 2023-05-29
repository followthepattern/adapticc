package sqlexpectations

import (
	"fmt"

	"github.com/followthepattern/adapticc/pkg/models"

	"github.com/DATA-DOG/go-sqlmock"
)

func ExpectGetUserByEmail(mock sqlmock.Sqlmock, result models.User, email string) {
	sqlQuery := fmt.Sprintf(`SELECT "active", "email", "first_name", "id", "last_name", "password", "registered_at", "salt" FROM "usr"."users" WHERE \("email" = '%v'\) LIMIT 1`, email)

	mock.ExpectQuery(sqlQuery).
		WillReturnRows(ModelToSQLMockRows(result))
}

func ExpectGetUserByID(mock sqlmock.Sqlmock, result models.User, id string) {
	sqlQuery := fmt.Sprintf(`SELECT "active", "email", "first_name", "id", "last_name", "password", "registered_at", "salt" FROM "usr"."users" WHERE \("id" = '%v'\) LIMIT 1`, id)

	mock.ExpectQuery(sqlQuery).
		WillReturnRows(ModelToSQLMockRows(result))
}

func ExpectUsers(mock sqlmock.Sqlmock, result []models.User, search string) {
	countQuery := fmt.Sprintf(`SELECT COUNT\(\*\) AS "count" FROM "usr"."users" WHERE \(\("first_name" LIKE '%%%v%%'\) OR \("last_name" LIKE '%%%v%%'\) OR \("email" LIKE '%%%v%%'\)\) LIMIT 1`, search, search, search)

	mock.ExpectQuery(countQuery).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(len(result)))

	sqlQuery := fmt.Sprintf(`SELECT "active", "email", "first_name", "id", "last_name", "password", "registered_at", "salt" FROM "usr"."users" WHERE \(\("first_name" LIKE '%%%v%%'\) OR \("last_name" LIKE '%%%v%%'\) OR \("email" LIKE '%%%v%%'\)\)`, search, search, search)

	SQLMockRows := ModelToSQLMockRows(result)
	mock.ExpectQuery(sqlQuery).
		WillReturnRows(SQLMockRows)
}

func ExpectCreateUser(mock sqlmock.Sqlmock, insert models.User) {
	sqlQuery := fmt.Sprintf(`INSERT INTO "usr"."users" \("active", "email", "first_name", "id", "last_name", "password", "registered_at", "salt"\) VALUES \(TRUE, '%v', '%v', '.*', '%v', '.*', NULL, '.*'\)`,
		*insert.Email,
		*insert.FirstName,
		*insert.LastName,
	)

	mock.ExpectExec(sqlQuery).
		WillReturnResult(sqlmock.NewResult(1, 1))
}

func ExpectUpdateUser(mock sqlmock.Sqlmock, update models.User) {
	sqlQuery := fmt.Sprintf(`UPDATE "usr"."users" SET "first_name"='%v',"last_name"='%v' WHERE \("id" = '%v'\)`,
		*update.FirstName,
		*update.LastName,
		*update.ID,
	)

	mock.ExpectExec(sqlQuery).
		WillReturnResult(sqlmock.NewResult(1, 1))
}
