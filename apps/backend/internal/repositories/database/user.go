package repositories

import (
	"backend/internal/container"
	"backend/internal/models"
	"backend/internal/utils"
	"errors"
	"fmt"

	"github.com/doug-martin/goqu/v9"
)

type User struct {
	db goqu.Database
}

func (User) tableName() string {
	return "users"
}

func ResolveUser(cont container.IContainer) (*User, error) {
	dependency := (*User)(nil)
	key := utils.GetKey(dependency)
	obj, err := cont.Resolve(key)

	if err != nil {
		return nil, err
	}

	if result, ok := obj.(*User); ok {
		return result, nil
	}

	return nil, fmt.Errorf("can't resolve %T", dependency)
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

func (repo User) Get(request models.ListRequest) (*models.ListResponse, error) {
	users := []models.User{}

	query := repo.db.From(repo.tableName()).Where(request.GetFilter())

	count, err := query.Count()
	if err != nil {
		return nil, err
	}

	err = query.ScanStructs(&users)
	if err != nil {
		return nil, err
	}
	return &models.ListResponse{
		Data:  users,
		Count: count,
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
		Where(goqu.Ex{"token": token})

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
