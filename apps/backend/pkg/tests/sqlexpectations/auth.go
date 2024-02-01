package sqlexpectations

import (
	"database/sql/driver"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/types"
)

func ExpectGetAuthUserByEmail(mock sqlmock.Sqlmock, result models.AuthUser, email types.String) {
	sqlQuery := fmt.Sprintf(`SELECT "active", "created_at", "creation_user_id", "email", "first_name", "id", "last_name", "password_hash", "salt", "update_user_id", "updated_at" FROM "usr"."users" WHERE ("email" = '%v') LIMIT 1`, email)

	columns := []string{
		"active", "created_at", "creation_user_id", "email", "first_name", "id", "last_name", "password_hash", "salt", "update_user_id", "updated_at",
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
		result.Salt,
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

func ExpectCreateAuthUser(mock sqlmock.Sqlmock, insert models.AuthUser) {
	sqlQuery := fmt.Sprintf(`
	INSERT INTO
		"usr"."users" ("active",
		"created_at",
		"email",
		"first_name",
		"id",
		"last_name",
		"password_hash",
		"salt")
	VALUES (FALSE, '.*', '%s', '%s', '.*', '%s', '.*', '.*')`,
		insert.Email,
		insert.FirstName,
		insert.LastName,
	)

	mock.ExpectExec(sqlQuery).
		WillReturnResult(sqlmock.NewResult(1, 1))
}
