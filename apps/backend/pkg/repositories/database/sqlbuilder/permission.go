package sqlbuilder

import (
	. "github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

func getRoleResourcePermissions(userID string) *SelectDataset {
	return From(S("usr").Table("user_role").As("ur")).
		LeftJoin(
			S("usr").Table("roles").As("r"),
			On(Ex{"ur.role_id": I("r.id")}),
		).
		LeftJoin(
			S("usr").Table("role_resource_permissions").As("rrp"),
			On(Ex{"rrp.role_id": I("ur.role_id")}),
		).
		Select("rrp.resource_id", "rrp.permission").
		Where(
			T("rrp").Col("permission").IsNotNull(),
			T("ur").Col("user_id").Eq(userID),
		)
}

func getUserResourcePermissions(userID string) *SelectDataset {
	return From(S("usr").Table("user_resource_permissions").As("urp")).
		Select("urp.resource_id", "urp.permission").
		Where(
			T("urp").Col("permission").IsNotNull(),
			T("urp").Col("user_id").Eq(userID),
		)
}

func getMergedResourcePermissions(userID string) *SelectDataset {
	rrpQuery := getRoleResourcePermissions(userID)
	urpQuery := getUserResourcePermissions(userID)

	query := From(rrpQuery.As("rp")).
		Select(
			L("COALESCE(rp.resource_id, up.resource_id)").As("resource_id"),
			L("COALESCE(rp.permission, 0) | COALESCE(up.permission, 0)").As("permissions"),
		).
		FullJoin(
			urpQuery.As("up"),
			On(Ex{"rp.resource_id": I("up.resource_id")}),
		)

	return query
}

func GetSelectWithPermissions(selection *SelectDataset, resourceName string, joinKey exp.IdentifierExpression, userID string) *SelectDataset {
	return selection.Join(
		getMergedResourcePermissions(userID).As("merged_resource_permissions"),
		On(
			Or(
				Ex{"merged_resource_permissions.resource_id": joinKey},
				Ex{"merged_resource_permissions.resource_id": resourceName},
			))).
		Where(L("merged_resource_permissions.permissions & 2").Gt(0))
}
