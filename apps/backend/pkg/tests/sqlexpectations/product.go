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
			"product_id",
			"title"
		FROM
			"usr"."products"
		INNER JOIN \(SELECT
				COALESCE\(rp.resource_id, up.resource_id\) AS "resource_id",
				COALESCE\(rp.permission, 0\) \| COALESCE\(up.permission, 0\) AS "permissions"
			FROM
				\(SELECT
					"rrp"."resource_id",
					"rrp"."permission"
				FROM
					"usr"."user_role" AS "ur"
				LEFT JOIN "usr"."roles" AS "r" ON
					\("ur"."role_id" = "r"."id"\)
				LEFT JOIN "usr"."role_resource_permission" AS "rrp" ON
					\("rrp"."role_id" = "ur"."role_id"\)
				WHERE
					\(\("rrp"."permission" IS NOT NULL\)
						AND \("ur"."user_id" = '%s'\)\)\) AS "rp"
			FULL JOIN \(SELECT
					"urp"."resource_id",
					"urp"."permission"
				FROM
					"usr"."user_resource_permission" AS "urp"
				WHERE
					\(\("urp"."permission" IS NOT NULL\)
						AND \("urp"."user_id" = '%s'\)\)\) AS "up" ON
				\("rp"."resource_id" = "up"."resource_id"\)\) AS "merged_resource_permissions" ON
			\(\("merged_resource_permissions"."resource_id" = "products"."product_id"\)
				OR \("merged_resource_permissions"."resource_id" = 'product'\)\)
		WHERE
			\(\("product_id" = '%s'\)
				AND \(merged_resource_permissions.permissions & 2 > 0\)\)
		LIMIT 1`,
		userID,
		userID,
		*result.ProductID)

	mock.ExpectQuery(sqlQuery).
		WillReturnRows(ModelToSQLMockRows(result))
}

func ExpectProducts(mock sqlmock.Sqlmock, userID string, page int, pageSize int, result []models.Product) {
	countQuery := fmt.Sprintf(
		`SELECT
		COUNT\(DISTINCT "products"."product_id"\)
	FROM
		"usr"."products"
	INNER JOIN \(
		SELECT
			COALESCE\(rp.resource_id, up.resource_id\) AS "resource_id",
			COALESCE\(rp.permission, 0\) | COALESCE\(up.permission, 0\) AS "permissions"
		FROM
			\(SELECT
				"rrp"."resource_id",
				"rrp"."permission"
			FROM
				"usr"."user_role" AS "ur"
			LEFT JOIN "usr"."roles" AS "r" ON
				\("ur"."role_id" = "r"."id"\)
			LEFT JOIN "usr"."role_resource_permission" AS "rrp" ON
				\("rrp"."role_id" = "ur"."role_id"\)
			WHERE
				\(\("rrp"."permission" IS NOT NULL\)
					AND \("ur"."user_id" = '%s'\)\)\) AS "rp"
		FULL JOIN \(SELECT
				"urp"."resource_id",
				"urp"."permission"
			FROM
				"usr"."user_resource_permission" AS "urp"
			WHERE
				\(\("urp"."permission" IS NOT NULL\)
					AND \("urp"."user_id" = '%s'\)\)\) AS "up" ON
			\("rp"."resource_id" = "up"."resource_id"\)\) AS "merged_resource_permissions" ON
		\(\("merged_resource_permissions"."resource_id" = "products"."product_id"\)
			OR \("merged_resource_permissions"."resource_id" = 'product'\)\)
	WHERE
		\(merged_resource_permissions.permissions & 2 > 0\)
	LIMIT 1`,
		userID,
		userID)

	mock.ExpectQuery(countQuery).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(len(result)))

	if page > 0 {
		page--
	}

	sqlQuery := fmt.Sprintf(`
	SELECT DISTINCT
		"description",
		"product_id",
		"title"
	FROM
		"usr"."products"
	INNER JOIN \(SELECT
			COALESCE\(rp.resource_id, up.resource_id\) AS "resource_id",
			COALESCE\(rp.permission, 0\) \| COALESCE\(up.permission, 0\) AS "permissions"
		FROM
			\(SELECT
				"rrp"."resource_id",
				"rrp"."permission"
			FROM
				"usr"."user_role" AS "ur"
			LEFT JOIN "usr"."roles" AS "r" ON
				\("ur"."role_id" = "r"."id"\)
			LEFT JOIN "usr"."role_resource_permission" AS "rrp" ON
				\("rrp"."role_id" = "ur"."role_id"\)
			WHERE
				\(\("rrp"."permission" IS NOT NULL\)
					AND \("ur"."user_id" = '%s'\)\)\) AS "rp"
		FULL JOIN \(SELECT
				"urp"."resource_id",
				"urp"."permission"
			FROM
				"usr"."user_resource_permission" AS "urp"
			WHERE
				\(\("urp"."permission" IS NOT NULL\)
					AND \("urp"."user_id" = '%s'\)\)\) AS "up" ON
			\("rp"."resource_id" = "up"."resource_id"\)\) AS "merged_resource_permissions" ON
		\(\("merged_resource_permissions"."resource_id" = "products"."product_id"\)
			OR \("merged_resource_permissions"."resource_id" = 'product'\)\)
	WHERE
		\(merged_resource_permissions.permissions & 2 > 0\)
	LIMIT %v OFFSET %v`,
		userID,
		userID,
		page*pageSize,
		pageSize,
	)

	mock.ExpectQuery(sqlQuery).
		WillReturnRows(ModelToSQLMockRows(result))
}

func CreateProduct(mock sqlmock.Sqlmock, userID string, product models.Product) {
	countQuery := fmt.Sprintf(`
	SELECT
		COUNT\(\*\) AS "count"
	FROM \(SELECT
			COALESCE\(rp.permission, 0\) \| COALESCE\(up.permission, 0\) AS "permissions"
		FROM
			\(SELECT
				"rrp"."resource_id",
				"rrp"."permission"
			FROM
				"usr"."user_role" AS "ur"
			LEFT JOIN "usr"."roles" AS "r" ON
				\("ur"."role_id" = "r"."id"\)
			LEFT JOIN "usr"."role_resource_permission" AS "rrp" ON
				\("rrp"."role_id" = "ur"."role_id"\)
			WHERE
				\(\("rrp"."permission" IS NOT NULL\)
					AND \("rrp"."resource_id" = 'product'\)
						AND \("ur"."user_id" = '%s'\)\)\) AS "rp"
		FULL JOIN \(SELECT
				"urp"."resource_id",
				"urp"."permission"
			FROM
				"usr"."user_resource_permission" AS "urp"
			WHERE
				\(\("urp"."permission" IS NOT NULL\)
					AND \("urp"."user_id" = '%s'\)
						AND \("urp"."resource_id" = 'product'\)\)\) AS "up" ON
			\("rp"."resource_id" = "up"."resource_id"\)\) AS "res"
	WHERE
	\(res.permissions & 1 > 0\)
	LIMIT 1`,
		userID,
		userID)

	mock.ExpectQuery(countQuery).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(1))

	sqlQuery := fmt.Sprintf(`INSERT INTO "usr"."products" \("description", "product_id", "title"\) VALUES \('%s', NULL, '%s'\)`,
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
		"description"='%s',"product_id"='%s',"title"='%s'
	FROM
		\(SELECT
			COALESCE\(rp.resource_id,
			up.resource_id\) AS "resource_id",
			COALESCE\(rp.permission,
			0\) \| COALESCE\(up.permission,
			0\) AS "permissions"
		FROM
			\(SELECT
				"rrp"."resource_id",
				"rrp"."permission"
			FROM
				"usr"."user_role" AS "ur"
			LEFT JOIN "usr"."roles" AS "r" ON
				\("ur"."role_id" = "r"."id"\)
			LEFT JOIN "usr"."role_resource_permission" AS "rrp" ON
				\("rrp"."role_id" = "ur"."role_id"\)
			WHERE
				\(\("rrp"."permission" IS NOT NULL\)
					AND \("ur"."user_id" = '%s'\)\)\) AS "rp"
		FULL JOIN \(SELECT
				"urp"."resource_id",
				"urp"."permission"
			FROM
				"usr"."user_resource_permission" AS "urp"
			WHERE
				\(\("urp"."permission" IS NOT NULL\)
					AND \("urp"."user_id" = '%s'\)\)\) AS "up" ON
			\("rp"."resource_id" = "up"."resource_id"\)\) AS "merged_resource_permissions"
	WHERE
		\(\(merged_resource_permissions.permissions & 4 > 0\)
			AND \(\("merged_resource_permissions"."resource_id" = "products"."product_id"\)
				OR \("merged_resource_permissions"."resource_id" = 'product'\)\)
				AND \("product_id" = '%s'\)\)`,
		*product.Description,
		*product.ProductID,
		*product.Title,
		userID,
		userID,
		*product.ProductID)

	mock.ExpectExec(sqlQuery).
		WillReturnResult(sqlmock.NewResult(1, 1))
}

func DeleteProduct(mock sqlmock.Sqlmock, userID string, product models.Product) {
	sqlQuery := fmt.Sprintf(`
	DELETE FROM
		"usr"."products"
	WHERE
		\(\("product_id" IN \(SELECT
			"products"."product_id"
		FROM
			"products"
		INNER JOIN \(SELECT
				COALESCE\(rp.resource_id,
				up.resource_id\) AS "resource_id",
				COALESCE\(rp.permission,
				0\) \| COALESCE\(up.permission,
				0\) AS "permissions"
			FROM
				\(SELECT
					"rrp"."resource_id",
					"rrp"."permission"
				FROM
					"usr"."user_role" AS "ur"
				LEFT JOIN "usr"."roles" AS "r" ON
					\("ur"."role_id" = "r"."id"\)
				LEFT JOIN "usr"."role_resource_permission" AS "rrp" ON
					\("rrp"."role_id" = "ur"."role_id"\)
				WHERE
					\(\("rrp"."permission" IS NOT NULL\)
						AND \("ur"."user_id" = '%s'\)\)\) AS "rp"
			FULL JOIN \(SELECT
					"urp"."resource_id",
					"urp"."permission"
				FROM
					"usr"."user_resource_permission" AS "urp"
				WHERE
					\(\("urp"."permission" IS NOT NULL\)
						AND \("urp"."user_id" = '%s'\)\)\) AS "up" ON
				\("rp"."resource_id" = "up"."resource_id"\)\) AS "merged_resource_permissions" ON
			\(\("merged_resource_permissions"."resource_id" = "products"."product_id"\)
				OR \("merged_resource_permissions"."resource_id" = 'product'\)\)
		WHERE
			\(merged_resource_permissions.permissions & 8 > 0\)\)\)
			AND \("product_id" = '%s'\)\)`,
		userID,
		userID,
		*product.ProductID)

	mock.ExpectExec(sqlQuery).
		WillReturnResult(sqlmock.NewResult(1, 1))
}
