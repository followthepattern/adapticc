package repositories

import (
	"backend/internal/container"
	"backend/internal/models"
	"backend/internal/utils"
	"errors"
	"fmt"

	"github.com/doug-martin/goqu/v9"
)

type UserToken struct {
	db goqu.Database
}

func (UserToken) tableName() string {
	return "user_tokens"
}

func ResolveUserToken(cont container.IContainer) (*UserToken, error) {
	dependency := (*UserToken)(nil)
	key := utils.GetKey(dependency)
	obj, err := cont.Resolve(key)

	if err != nil {
		return nil, err
	}

	if result, ok := obj.(*UserToken); ok {
		return result, nil
	}

	return nil, fmt.Errorf("can't resolve %T", dependency)
}

func UserTokenDependencyConstructor(cont container.IContainer) (interface{}, error) {
	db := goqu.New("postgres", cont.GetDB())

	if db == nil {
		return nil, errors.New("db is null")
	}

	return &UserToken{db: *db}, nil
}

func (repo UserToken) Create(userToken *models.UserToken) (err error) {
	_, err = repo.db.Insert(repo.tableName()).Rows(*userToken).Executor().Exec()
	return
}

func (repo UserToken) GetByToken(token string) (*models.User, error) {
	user := models.User{}

	query := repo.db.From(repo.tableName()).Where(goqu.Ex{"token": token})

	_, err := query.ScanStruct(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (dao UserToken) Update(user *models.UserToken) error {
	return nil
}

func (dao UserToken) Delete(id string) error {
	return nil
}
