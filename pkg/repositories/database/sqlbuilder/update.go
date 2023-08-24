package sqlbuilder

import (
	. "github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

func GetUpdateWithPermissions(ds *UpdateDataset, resourceName string, joinKey exp.IdentifierExpression, userID string) *UpdateDataset {
	return ds.From(getMergedResourcePermissions(userID).As("merged_resource_permissions")).
		Where(L("merged_resource_permissions.permissions & 4").Gt(0),
			Or(
				Ex{"merged_resource_permissions.resource_id": joinKey},
				Ex{"merged_resource_permissions.resource_id": resourceName},
			))
}
