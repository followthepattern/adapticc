package sqlbuilder

import (
	"github.com/followthepattern/adapticc/models"
	"github.com/followthepattern/adapticc/types"
	. "github.com/followthepattern/goqu/v9"
	"github.com/followthepattern/goqu/v9/exp"
)

func WithPagination(query *SelectDataset, pagination models.Pagination) *SelectDataset {
	if pagination.Page.Data < 1 {
		pagination.Page = types.UintFrom(models.DefaultPage)
	}

	if pagination.PageSize.Data > 0 {
		offset := (pagination.Page.Data - 1) * pagination.PageSize.Data

		query = query.Offset(offset)
		query = query.Limit(pagination.PageSize.Data)
	}

	return query
}

func WithOrderBy(query *SelectDataset, orderBy []models.OrderBy) *SelectDataset {
	orderLength := len(orderBy)
	if orderLength > 0 {
		orderExpressions := make([]exp.OrderedExpression, orderLength)
		for i, order := range orderBy {
			orderExpressions[i] = I(order.Name).Asc()
			if order.Desc.IsValid() && order.Desc.Data {
				orderExpressions[i] = I(order.Name).Desc()
			}
		}
		query = query.Order(orderExpressions...)
	}

	return query
}
