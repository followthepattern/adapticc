package sqlbuilder

import (
	. "github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

func GetDeleteWithPermissions(ds *DeleteDataset, resourceName string, joinKey exp.IdentifierExpression, userID string) *DeleteDataset {
	col := joinKey.GetCol()

	colStr := col.(string)

	ds = ds.Where(
		Ex{colStr: From(joinKey.GetTable()).Select(joinKey).
			Join(
				getMergedResourcePermissions(userID).As("merged_resource_permissions"),
				On(
					Or(
						Ex{"merged_resource_permissions.resource_id": joinKey},
						Ex{"merged_resource_permissions.resource_id": resourceName},
					))).
			Where(L("merged_resource_permissions.permissions & 8").Gt(0))})

	return ds
}
