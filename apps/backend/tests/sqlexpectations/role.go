package sqlexpectations

import (
	"database/sql/driver"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/followthepattern/adapticc/features/role"
	"github.com/followthepattern/adapticc/types"
)

var roleColumns = []string{
	"id",
	"code",
	"name",
}

func ExpectRolesByUserID(mock sqlmock.Sqlmock, results []role.RoleModel, userID types.String) {
	sqlQuery := fmt.Sprintf(`
	SELECT
		"r".\*
	FROM
		"usr"."user_role" AS "ur"
	INNER JOIN
		"usr"."roles" AS "r" ON ("r"."id" = "ur"."role_id")
	WHERE ("user_id" = '%v')`, userID)

	rows := sqlmock.NewRows(roleColumns)

	for _, result := range results {
		values := []driver.Value{
			result.ID,
			result.CreatedAt,
			result.CreationUserID,
		}
		rows.AddRow(values...)
	}

	mock.ExpectQuery(sqlQuery).
		WillReturnRows(rows)
}
