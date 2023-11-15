package sqlexpectations

import (
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/types"
)

func ExpectRolesByUserID(mock sqlmock.Sqlmock, result []models.Role, userID types.String) {
	sqlQuery := fmt.Sprintf(`
	SELECT
		"r".\*
	FROM
		"usr"."user_role" AS "ur"
	INNER JOIN
		"usr"."roles" AS "r" ON ("r"."id" = "ur"."role_id")
	WHERE ("user_id" = '%v')`, userID)

	mock.ExpectQuery(sqlQuery).
		WillReturnRows(ModelToSQLMockRows(result))
}
