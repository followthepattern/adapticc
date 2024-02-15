package sqlexpectations

import (
	"database/sql/driver"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/followthepattern/adapticc/features/auth"
	"github.com/followthepattern/adapticc/types"
)

func ExpectGetAuthUserByEmail(mock sqlmock.Sqlmock, result auth.AuthUser, email types.String) {
	sqlQuery := fmt.Sprintf(`SELECT "active", "created_at", "creation_user_id", "email", "first_name", "id", "last_name", "password_hash", "update_user_id", "updated_at" FROM "usr"."users" WHERE ("email" = '%v') LIMIT 1`, email)

	columns := []string{
		"active", "created_at", "creation_user_id", "email", "first_name", "id", "last_name", "password_hash", "update_user_id", "updated_at",
	}

	rows := sqlmock.NewRows(columns)

	values := []driver.Value{
		result.Active,
		result.CreatedAt,
		result.CreationUserID,
		result.Email,
		result.FirstName,
		result.ID,
		result.LastName,
		result.PasswordHash,
		result.UpdateUserID,
		result.UpdatedAt,
	}

	rows.AddRow(values...)

	mock.ExpectQuery(sqlQuery).
		WillReturnRows(rows)
}

func ExpectVerifyEmail(mock sqlmock.Sqlmock, count int, email types.String) {
	sqlQuery := fmt.Sprintf(`SELECT COUNT(\*) AS "count" FROM "usr"."users" WHERE ("email" = '%s') LIMIT 1`, email)

	mock.ExpectQuery(sqlQuery).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(count))
}

func ExpectCreateAuthUser(mock sqlmock.Sqlmock, insert auth.AuthUser) {
	sqlQuery := fmt.Sprintf(`
	INSERT INTO
		"usr"."users" ("active",
		"created_at",
		"email",
		"first_name",
		"id",
		"last_name",
		"password_hash")
	VALUES (FALSE, '.*', '%s', '%s', '.*', '%s', '.*')`,
		insert.Email,
		insert.FirstName,
		insert.LastName,
	)

	mock.ExpectExec(sqlQuery).
		WillReturnResult(sqlmock.NewResult(1, 1))
}

func ExpectRoleIDsByUserID(mock sqlmock.Sqlmock, results []string, userID types.String) {
	sqlQuery := fmt.Sprintf(`
	SELECT
		"r"."code"
	FROM
		"usr"."user_role" AS "ur"
	INNER JOIN
		"usr"."roles" AS "r" ON ("r"."id" = "ur"."role_id")
	WHERE ("user_id" = '%v')`, userID)

	rows := sqlmock.NewRows([]string{"code"})

	for _, result := range results {
		values := []driver.Value{
			result,
		}
		rows.AddRow(values...)
	}

	mock.ExpectQuery(sqlQuery).
		WillReturnRows(rows)
}
