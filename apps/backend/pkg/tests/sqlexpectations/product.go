package sqlexpectations

import (
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/followthepattern/adapticc/pkg/models"
)

func ExpectProduct(mock sqlmock.Sqlmock, userID string, result models.Product) {
	sqlQuery := fmt.Sprintf(
		`SELECT
			"description",
			"id",
			"title"
		FROM
			"usr"."products"
		INNER JOIN (SELECT
				COALESCE(rp.resource_id, up.resource_id) AS "resource_id",
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
			(("merged_resource_permissions"."resource_id" = "products"."id")
				OR ("merged_resource_permissions"."resource_id" = 'PRODUCT'))
		WHERE
			(("id" = '%s')
				AND (merged_resource_permissions.permissions & 2 > 0))
		LIMIT 1`,
		userID,
		userID,
		*result.ID)

	mock.ExpectQuery(sqlQuery).
		WillReturnRows(ModelToSQLMockRows(result))
}

func ExpectProducts(mock sqlmock.Sqlmock, userID string, filter models.ListFilter, page int, pageSize int, result []models.Product) {
	countQuery := fmt.Sprintf(`
	SELECT
		COUNT(DISTINCT "products"."id")
	FROM
		"usr"."products"
	INNER JOIN (SELECT
			COALESCE(rp.resource_id, up.resource_id) AS "resource_id",
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
		(("merged_resource_permissions"."resource_id" = "products"."id")
			OR ("merged_resource_permissions"."resource_id" = 'PRODUCT'))
	WHERE ((("id" LIKE '%%%s%%') OR
			("title" LIKE '%%%s%%')) AND
		(merged_resource_permissions.permissions & 2 > 0))
	LIMIT 1`,
		userID,
		userID,
		*filter.Search,
		*filter.Search,
	)

	mock.ExpectQuery(countQuery).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(len(result)))

	if page > 0 {
		page--
	}

	sqlQuery := fmt.Sprintf(`
	SELECT DISTINCT
		"description",
		"id",
		"title"
	FROM
		"usr"."products"
	INNER JOIN (SELECT
			COALESCE(rp.resource_id, up.resource_id) AS "resource_id",
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
		(("merged_resource_permissions"."resource_id" = "products"."id")
			OR ("merged_resource_permissions"."resource_id" = 'PRODUCT'))
	WHERE 
		((("id" LIKE '%%%s%%') OR
			("title" LIKE '%%%s%%')) AND
		(merged_resource_permissions.permissions & 2 > 0))
	LIMIT %v OFFSET %v`,
		userID,
		userID,
		*filter.Search,
		*filter.Search,
		page*pageSize,
		pageSize,
	)

	mock.ExpectQuery(sqlQuery).
		WillReturnRows(ModelToSQLMockRows(result))
}

func CreateProduct(mock sqlmock.Sqlmock, userID string, product models.Product) {
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
					AND ("rrp"."resource_id" = 'PRODUCT')
						AND ("ur"."user_id" = '%s'))) AS "rp"
		FULL JOIN (SELECT
				"urp"."resource_id",
				"urp"."permission"
			FROM
				"usr"."user_resource_permissions" AS "urp"
			WHERE
				(("urp"."permission" IS NOT NULL)
					AND ("urp"."user_id" = '%s')
						AND ("urp"."resource_id" = 'PRODUCT'))) AS "up" ON
			("rp"."resource_id" = "up"."resource_id")) AS "res"
	WHERE
	(res.permissions & 1 > 0)
	LIMIT 1`,
		userID,
		userID)

	mock.ExpectQuery(countQuery).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(1))

	sqlQuery := fmt.Sprintf(`INSERT INTO "usr"."products" ("description", "id", "title") VALUES ('%s', .*, '%s')`,
		*product.Description,
		*product.Title,
	)

	mock.ExpectExec(sqlQuery).
		WillReturnResult(sqlmock.NewResult(1, 1))
}

func UpdateProduct(mock sqlmock.Sqlmock, userID string, product models.Product) {
	sqlQuery := fmt.Sprintf(`
	UPDATE
		"usr"."products"
	SET
		"description"='%s',"id"='%s',"title"='%s'
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
		((merged_resource_permissions.permissions & 4 > 0)
			AND (("merged_resource_permissions"."resource_id" = "products"."id")
				OR ("merged_resource_permissions"."resource_id" = 'PRODUCT'))
				AND ("id" = '%s'))`,
		*product.Description,
		*product.ID,
		*product.Title,
		userID,
		userID,
		*product.ID)

	mock.ExpectExec(sqlQuery).
		WillReturnResult(sqlmock.NewResult(1, 1))
}

func DeleteProduct(mock sqlmock.Sqlmock, userID string, product models.Product) {
	sqlQuery := fmt.Sprintf(`
	DELETE FROM
		"usr"."products"
	WHERE
		(("id" IN (SELECT
			"usr"."products"."id"
		FROM
			"usr"."products"
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
			(("merged_resource_permissions"."resource_id" = "usr"."products"."id")
				OR ("merged_resource_permissions"."resource_id" = 'PRODUCT'))
		WHERE
			(merged_resource_permissions.permissions & 8 > 0)))
			AND ("id" = '%s'))`,
		userID,
		userID,
		*product.ID)

	mock.ExpectExec(sqlQuery).
		WillReturnResult(sqlmock.NewResult(1, 1))
}
