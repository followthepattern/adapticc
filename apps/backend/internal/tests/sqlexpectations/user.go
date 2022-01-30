package sqlexpectations

import (
	"backend/internal/models"
	"backend/internal/tests/data_generator"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
)

func ExpectGetUserByEmail(mock sqlmock.Sqlmock, email string, result models.User) {
	sqlQuery := fmt.Sprintf(`SELECT "active", "created_at", "creation_user_id", "email", "first_name", "id", "last_login_at", "last_name", "password_hash", "salt", "update_user_id", "updated_at" FROM "users" WHERE \("email" = '%v'\) LIMIT 1`, email)

	mock.ExpectQuery(sqlQuery).
		WillReturnRows(data_generator.ModelToSQLMockRows(result))
}
