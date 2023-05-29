package sqlbuilder

import (
	. "github.com/doug-martin/goqu/v9"
)

func getInsertRoleResourcePermissions(userID string, resourceName string) *SelectDataset {
	return From(S("usr").Table("user_role").As("ur")).
		LeftJoin(
			S("usr").Table("roles").As("r"),
			On(Ex{"ur.role_id": I("r.id")}),
		).
		LeftJoin(
			S("usr").Table("role_resource_permission").As("rrp"),
			On(Ex{"rrp.role_id": I("ur.role_id")}),
		).
		Select("rrp.resource_id", "rrp.permission").
		Where(
			T("rrp").Col("permission").IsNotNull(),
			T("rrp").Col("resource_id").Eq(resourceName),
			T("ur").Col("user_id").Eq(userID),
		)
}

func getInsertUserResourcePermissions(userID string, resourceName string) *SelectDataset {
	return From(S("usr").Table("user_resource_permission").As("urp")).
		Select("urp.resource_id", "urp.permission").
		Where(
			T("urp").Col("permission").IsNotNull(),
			T("urp").Col("user_id").Eq(userID),
			T("urp").Col("resource_id").Eq(resourceName),
		)
}

func getInsertMergedResourcePermissions(userID string, resourceName string) *SelectDataset {
	rrpQuery := getInsertRoleResourcePermissions(userID, resourceName)
	urpQuery := getInsertUserResourcePermissions(userID, resourceName)

	query := From(rrpQuery.As("rp")).
		Select(
			L("COALESCE(rp.permission, 0) | COALESCE(up.permission, 0)").As("permissions"),
		).
		FullJoin(
			urpQuery.As("up"),
			On(Ex{"rp.resource_id": I("up.resource_id")}),
		)

	return query
}

func GetInsertWithPermissions(db Database, resourceName string, userID string) (int64, error) {
	return db.From(getInsertMergedResourcePermissions(userID, resourceName).As("res")).Where(L("res.permissions & 1").Gt(0)).Count()
}
