package sqlexpectations

import (
	"backend/internal/models"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
)

func ExpectInsertToken(mock sqlmock.Sqlmock, result models.User) {
	sqlQuery := fmt.Sprintf(`INSERT INTO "user_tokens" \("expires_at", "token", "user_id"\) VALUES \('.*', '.*', '%v'\)`, *result.ID)

	mock.ExpectExec(sqlQuery).WillReturnResult(sqlmock.NewResult(1, 1))
}
