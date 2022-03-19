package repositories

import (
	"backend/internal/container"
	"backend/internal/models"
	"backend/internal/utils"
	"errors"
	"fmt"
	"time"

	"github.com/doug-martin/goqu/v9"
)

type User struct {
	db goqu.Database
}

func (User) tableName() string {
	return "users"
}

func ResolveUser(cont container.IContainer) (dependency *User, err error) {
	key := utils.GetKey(dependency)
	obj, err := cont.Resolve(key)
	if err != nil {
		return nil, fmt.Errorf("can't resolve %T, error: %s", dependency, err.Error())
	}

	dependency, ok := obj.(*User)
	if !ok {
		return nil, fmt.Errorf("%T can't be resolved to %T", obj, dependency)
	}

	return dependency, nil
}

func UserDependencyConstructor(cont container.IContainer) (interface{}, error) {
	db := goqu.New("postgres", cont.GetDB())

	if db == nil {
		return nil, errors.New("db is null")
	}

	return &User{db: *db}, nil
}

func (repo User) Create(user *models.User) (err error) {
	_, err = repo.db.Insert(repo.tableName()).Rows(*user).Executor().Exec()
	return
}

func (repo User) Get(request models.UserListRequest) (*models.UserListResponse, error) {
	users := []models.User{}

	query := repo.db.From(repo.tableName())

	if request.Email != nil {
		query = query.Where(goqu.Ex{"email": request.Email})
	}

	if request.Name != nil {
		pattern := fmt.Sprintf("%%%v%%", *request.Name)
		query = query.Where(
			goqu.Or(
				goqu.I("first_name").Like(pattern),
				goqu.I("last_name").Like(pattern),
			))
	}

	count, err := query.Count()
	if err != nil {
		return nil, err
	}

	if request.PageSize != nil {
		query = query.Limit(*request.PageSize)
	}

	if request.Page != nil && request.PageSize != nil {
		page := *request.Page
		if page > 0 {
			page--
		}
		query = query.Offset(page * *request.PageSize)
	}

	err = query.ScanStructs(&users)
	if err != nil {
		return nil, err
	}
	return &models.UserListResponse{
		Data: users,
		ListResponse: models.ListResponse{
			Count:    count,
			Page:     request.Page,
			PageSize: request.PageSize,
		},
	}, nil
}

func (repo User) GetByID(id string) (*models.User, error) {
	user := models.User{}

	query := repo.db.From(repo.tableName()).Where(goqu.Ex{"id": id})

	_, err := query.ScanStruct(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (repo User) GetByToken(token string) (*models.User, error) {
	user := models.User{}

	query := repo.db.From(repo.tableName()).
		LeftJoin(
			goqu.T("user_tokens"),
			goqu.On(goqu.Ex{
				repo.tableName() + ".id": goqu.I("user_tokens.user_id"),
			}),
		).
		Where(goqu.Ex{"token": token}, goqu.C("expires_at").Gt(time.Now()))
	_, err := query.ScanStruct(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (repo User) GetByEmail(email string) (*models.User, error) {
	user := models.User{}

	query := repo.db.From(repo.tableName()).Where(goqu.Ex{"email": email})

	_, err := query.ScanStruct(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (dao User) Update(user *models.User) error {
	return nil
}

func (dao User) Delete(id string) error {
	return nil
}
