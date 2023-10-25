package database

import (
	"context"
	"database/sql"
	"fmt"

	. "github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/utils/pointers"
)

var (
	roleTableName     = S("usr").Table("roles")
	userRoleTableName = S("usr").Table("user_role")
)

type Role struct {
	db  *Database
	ctx context.Context
}

func NewRole(ctx context.Context, database *sql.DB) Role {
	db := New("postgres", database)

	return Role{
		ctx: ctx,
		db:  db,
	}
}

func (repo Role) GetByID(id string) (*models.Role, error) {
	var data models.Role

	_, err := repo.db.From(roleTableName).
		Where(Ex{"id": id}).
		ScanStruct(&data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (repo Role) Get(request models.RoleListRequestParams) (*models.RoleListResponse, error) {
	data := []models.Role{}

	query := repo.db.From(roleTableName)

	if request.Filter.Search != nil {
		pattern := fmt.Sprintf("%%%s%%", *request.Filter.Search)
		query = query.Where(
			Or(
				I("id").Like(pattern),
				I("name").Like(pattern),
				I("code").Like(pattern),
			))
	}

	count, err := query.Count()
	if err != nil {
		return nil, err
	}

	if request.Pagination.Page == nil {
		request.Pagination.Page = pointers.ToPtr[uint](models.DefaultPage)
	}

	if request.Pagination.PageSize != nil {
		page := *request.Pagination.Page
		if page > 0 {
			page--
		}

		query = query.Offset(page * *request.Pagination.PageSize)
		query = query.Limit(*request.Pagination.PageSize)
	}

	orderLength := len(request.OrderBy)
	if orderLength > 0 {
		orderExpressions := make([]exp.OrderedExpression, orderLength)
		for i, order := range request.OrderBy {
			orderExpressions[i] = I(order.Name).Asc()
			if order.Desc != nil && *order.Desc {
				orderExpressions[i] = I(order.Name).Desc()
			}
		}
		query = query.Order(orderExpressions...)
	}

	err = query.ScanStructs(&data)
	if err != nil {
		return nil, err
	}

	result := models.RoleListResponse{
		Count:    count,
		PageSize: request.Pagination.PageSize,
		Page:     request.Pagination.Page,
		Data:     data,
	}

	return &result, nil
}

func (repo Role) GetRolesByUserID(userID string) ([]models.Role, error) {
	var data []models.Role

	err := repo.db.From(userRoleTableName.As("ur")).
		Join(roleTableName.As("r"),
			On(Ex{"r.id": I("ur.role_id")})).
		Where(Ex{"user_id": userID}).
		ScanStructs(&data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (repo Role) GetRoleCodes(userID string) ([]string, error) {
	roles, err := repo.GetRolesByUserID(userID)
	if err != nil {
		return nil, err
	}

	result := make([]string, len(roles))

	for i, role := range roles {
		result[i] = role.Code
	}

	return result, nil
}
