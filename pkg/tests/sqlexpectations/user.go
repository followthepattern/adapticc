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

func ExpectUsers(mock sqlmock.Sqlmock, userID string, result []models.User, listRequestParams models.UserListRequestParams) {
	countQuery := fmt.Sprintf(`
	SELECT
		COUNT(\*) AS "count"
	FROM
		"usr"."users"
	INNER JOIN (
		SELECT
			COALESCE(rp.resource_id,
			up.resource_id) AS "resource_id",
			COALESCE(rp.permission,
			0) | COALESCE(up.permission,
			0) AS "permissions"
		FROM
			(SELECT
				"rrp"."resource_id",
				"rrp"."permission"
			FROM
				"usr"."user_role" AS "ur"
			LEFT JOIN "usr"."roles" AS "r" ON
				("ur"."role_id" = "r"."id")
			LEFT JOIN "usr"."role_resource_permissions" AS "rrp" ON
				("rrp"."role_id" = "ur"."role_id")
			WHERE
				(("rrp"."permission" IS NOT NULL)
					AND ("ur"."user_id" = '%s'))) AS "rp"
		FULL JOIN (SELECT
				"urp"."resource_id",
				"urp"."permission"
			FROM
				"usr"."user_resource_permissions" AS "urp"
			WHERE
				(("urp"."permission" IS NOT NULL)
					AND ("urp"."user_id" = '%s'))) AS "up" ON
			("rp"."resource_id" = "up"."resource_id")) AS "merged_resource_permissions" ON
		(("merged_resource_permissions"."resource_id" = "users"."id")
			OR ("merged_resource_permissions"."resource_id" = 'USER'))
	WHERE
		((("first_name" LIKE '%%%s%%')
			OR ("last_name" LIKE '%%%s%%')
				OR ("email" LIKE '%%%s%%'))
			AND (merged_resource_permissions.permissions & 2 > 0))
	LIMIT 1`,
		userID,
		userID,
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
	INNER JOIN (SELECT
			COALESCE(rp.resource_id,
			up.resource_id) AS "resource_id",
			COALESCE(rp.permission,
			0) | COALESCE(up.permission,
			0) AS "permissions"
		FROM
			(SELECT
				"rrp"."resource_id",
				"rrp"."permission"
			FROM
				"usr"."user_role" AS "ur"
			LEFT JOIN "usr"."roles" AS "r" ON
				("ur"."role_id" = "r"."id")
			LEFT JOIN "usr"."role_resource_permissions" AS "rrp" ON
				("rrp"."role_id" = "ur"."role_id")
			WHERE
				(("rrp"."permission" IS NOT NULL)
					AND ("ur"."user_id" = '%s'))) AS "rp"
		FULL JOIN (SELECT
				"urp"."resource_id",
				"urp"."permission"
			FROM
				"usr"."user_resource_permissions" AS "urp"
			WHERE
				(("urp"."permission" IS NOT NULL)
					AND ("urp"."user_id" = '%s'))) AS "up" ON
			("rp"."resource_id" = "up"."resource_id")) AS "merged_resource_permissions" ON
		(("merged_resource_permissions"."resource_id" = "users"."id")
			OR ("merged_resource_permissions"."resource_id" = 'USER'))
	WHERE
		((("first_name" LIKE '%%%s%%')
			OR ("last_name" LIKE '%%%s%%')
				OR ("email" LIKE '%%%s%%'))
			AND (merged_resource_permissions.permissions & 2 > 0))
	LIMIT %v`,
		userID,
		userID,
		*listRequestParams.Filter.Search,
		*listRequestParams.Filter.Search,
		*listRequestParams.Filter.Search,
		*listRequestParams.Pagination.PageSize)

	SQLMockRows := ModelToSQLMockRows(result)
	mock.ExpectQuery(sqlQuery).
		WillReturnRows(SQLMockRows)
}

func ExpectUsersWithoutPaging(mock sqlmock.Sqlmock, userID string, result []models.User, listRequestParams models.UserListRequestParams) {
	countQuery := fmt.Sprintf(`
	SELECT
		COUNT(\*) AS "count"
	FROM
		"usr"."users"
	INNER JOIN (SELECT
			COALESCE(rp.resource_id,
			up.resource_id) AS "resource_id",
			COALESCE(rp.permission,
			0) | COALESCE(up.permission,
			0) AS "permissions"
		FROM
			(SELECT
				"rrp"."resource_id",
				"rrp"."permission"
			FROM
				"usr"."user_role" AS "ur"
			LEFT JOIN "usr"."roles" AS "r" ON
				("ur"."role_id" = "r"."id")
			LEFT JOIN "usr"."role_resource_permissions" AS "rrp" ON
				("rrp"."role_id" = "ur"."role_id")
			WHERE
				(("rrp"."permission" IS NOT NULL)
					AND ("ur"."user_id" = '%s'))) AS "rp"
		FULL JOIN (SELECT
				"urp"."resource_id",
				"urp"."permission"
			FROM
				"usr"."user_resource_permissions" AS "urp"
			WHERE
				(("urp"."permission" IS NOT NULL)
					AND ("urp"."user_id" = '%s'))) AS "up" ON
			("rp"."resource_id" = "up"."resource_id")) AS "merged_resource_permissions" ON
		(("merged_resource_permissions"."resource_id" = "users"."id")
			OR ("merged_resource_permissions"."resource_id" = 'USER'))
	WHERE
		((("first_name" LIKE '%%%s%%')
			OR ("last_name" LIKE '%%%s%%')
				OR ("email" LIKE '%%%s%%'))
			AND (merged_resource_permissions.permissions & 2 > 0))
	LIMIT 1`,
		userID,
		userID,
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
	INNER JOIN (SELECT
			COALESCE(rp.resource_id,
			up.resource_id) AS "resource_id",
			COALESCE(rp.permission,
			0) | COALESCE(up.permission,
			0) AS "permissions"
		FROM
			(SELECT
				"rrp"."resource_id",
				"rrp"."permission"
			FROM
				"usr"."user_role" AS "ur"
			LEFT JOIN "usr"."roles" AS "r" ON
				("ur"."role_id" = "r"."id")
			LEFT JOIN "usr"."role_resource_permissions" AS "rrp" ON
				("rrp"."role_id" = "ur"."role_id")
			WHERE
				(("rrp"."permission" IS NOT NULL)
					AND ("ur"."user_id" = '%s'))) AS "rp"
		FULL JOIN (SELECT
				"urp"."resource_id",
				"urp"."permission"
			FROM
				"usr"."user_resource_permissions" AS "urp"
			WHERE
				(("urp"."permission" IS NOT NULL)
					AND ("urp"."user_id" = '%s'))) AS "up" ON
			("rp"."resource_id" = "up"."resource_id")) AS "merged_resource_permissions" ON
		(("merged_resource_permissions"."resource_id" = "users"."id")
			OR ("merged_resource_permissions"."resource_id" = 'USER'))
	WHERE
		((("first_name" LIKE '%%%s%%')
			OR ("last_name" LIKE '%%%s%%')
				OR ("email" LIKE '%%%s%%'))
			AND (merged_resource_permissions.permissions & 2 > 0))`,
		userID,
		userID,
		*listRequestParams.Filter.Search,
		*listRequestParams.Filter.Search,
		*listRequestParams.Filter.Search)

	SQLMockRows := ModelToSQLMockRows(result)
	mock.ExpectQuery(sqlQuery).
		WillReturnRows(SQLMockRows)
}

func ExpectCreateUser(mock sqlmock.Sqlmock, userID string, insert models.User) {
	countQuery := fmt.Sprintf(`
	SELECT
		COUNT(\*) AS "count"
	FROM (SELECT
			COALESCE(rp.permission, 0) | COALESCE(up.permission, 0) AS "permissions"
		FROM
			(SELECT
				"rrp"."resource_id",
				"rrp"."permission"
			FROM
				"usr"."user_role" AS "ur"
			LEFT JOIN "usr"."roles" AS "r" ON
				("ur"."role_id" = "r"."id")
			LEFT JOIN "usr"."role_resource_permissions" AS "rrp" ON
				("rrp"."role_id" = "ur"."role_id")
			WHERE
				(("rrp"."permission" IS NOT NULL)
					AND ("rrp"."resource_id" = 'USER')
						AND ("ur"."user_id" = '%s'))) AS "rp"
		FULL JOIN (SELECT
				"urp"."resource_id",
				"urp"."permission"
			FROM
				"usr"."user_resource_permissions" AS "urp"
			WHERE
				(("urp"."permission" IS NOT NULL)
					AND ("urp"."user_id" = '%s')
						AND ("urp"."resource_id" = 'USER'))) AS "up" ON
			("rp"."resource_id" = "up"."resource_id")) AS "res"
	WHERE
	(res.permissions & 1 > 0)
	LIMIT 1`,
		userID,
		userID)

	mock.ExpectQuery(countQuery).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(1))

	sqlQuery := fmt.Sprintf(`INSERT INTO "usr"."users" ("active", "created_at", "creation_user_id", "email", "first_name", "id", "last_name", "update_user_id", "updated_at") VALUES (FALSE, '.*', '%v', '%v', '%v', '.*', '%v', NULL, NULL)`,
		*insert.CreationUserID,
		*insert.Email,
		*insert.FirstName,
		*insert.LastName,
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
	FROM
		(SELECT
			COALESCE(rp.resource_id,
			up.resource_id) AS "resource_id",
			COALESCE(rp.permission,
			0) | COALESCE(up.permission,
			0) AS "permissions"
		FROM
			(SELECT
				"rrp"."resource_id",
				"rrp"."permission"
			FROM
				"usr"."user_role" AS "ur"
			LEFT JOIN "usr"."roles" AS "r" ON
				("ur"."role_id" = "r"."id")
			LEFT JOIN "usr"."role_resource_permissions" AS "rrp" ON
				("rrp"."role_id" = "ur"."role_id")
			WHERE
				(("rrp"."permission" IS NOT NULL)
					AND ("ur"."user_id" = '%s'))) AS "rp"
		FULL JOIN (SELECT
				"urp"."resource_id",
				"urp"."permission"
			FROM
				"usr"."user_resource_permissions" AS "urp"
			WHERE
				(("urp"."permission" IS NOT NULL)
					AND ("urp"."user_id" = '%s'))) AS "up" ON
			("rp"."resource_id" = "up"."resource_id")) AS "merged_resource_permissions"
	WHERE
		(("id" = '%s')
			AND (merged_resource_permissions.permissions & 4 > 0)
				AND (("merged_resource_permissions"."resource_id" = "users"."id")
					OR ("merged_resource_permissions"."resource_id" = 'USER')))`,
		*model.FirstName,
		*model.LastName,
		userID,
		userID,
		userID,
		*model.ID,
	)

	mock.ExpectExec(sqlQuery).
		WillReturnResult(sqlmock.NewResult(1, 1))
}

func ExpectDeleteUser(mock sqlmock.Sqlmock, userID string, user models.User) {
	sqlQuery := fmt.Sprintf(`
	DELETE FROM
		"usr"."users"
	WHERE
		(("id" = '%s') AND
		("id" IN (SELECT
			"usr"."users"."id"
		FROM
			"usr"."users"
		INNER JOIN (SELECT
				COALESCE(rp.resource_id,
				up.resource_id) AS "resource_id",
				COALESCE(rp.permission,
				0) | COALESCE(up.permission,
				0) AS "permissions"
			FROM
				(SELECT
					"rrp"."resource_id",
					"rrp"."permission"
				FROM
					"usr"."user_role" AS "ur"
				LEFT JOIN "usr"."roles" AS "r" ON
					("ur"."role_id" = "r"."id")
				LEFT JOIN "usr"."role_resource_permissions" AS "rrp" ON
					("rrp"."role_id" = "ur"."role_id")
				WHERE
					(("rrp"."permission" IS NOT NULL)
						AND ("ur"."user_id" = '%s'))) AS "rp"
			FULL JOIN (SELECT
					"urp"."resource_id",
					"urp"."permission"
				FROM
					"usr"."user_resource_permissions" AS "urp"
				WHERE
					(("urp"."permission" IS NOT NULL)
						AND ("urp"."user_id" = '%s'))) AS "up" ON
				("rp"."resource_id" = "up"."resource_id")) AS "merged_resource_permissions" ON
			(("merged_resource_permissions"."resource_id" = "usr"."users"."id")
				OR ("merged_resource_permissions"."resource_id" = 'USER'))
		WHERE
			(merged_resource_permissions.permissions & 8 > 0))))`,
		*user.ID,
		userID,
		userID)

	mock.ExpectExec(sqlQuery).
		WillReturnResult(sqlmock.NewResult(1, 1))
}
