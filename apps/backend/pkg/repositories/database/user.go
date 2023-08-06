package database

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/repositories/database/sqlbuilder"
	"github.com/followthepattern/adapticc/pkg/request"
	"github.com/followthepattern/adapticc/pkg/utils/pointers"

	. "github.com/doug-martin/goqu/v9"
)

type UserMsgChannel chan models.UserMsg

func RegisterUserChannel(cont *container.Container) {
	if cont == nil {
		return
	}
	requestChan := make(UserMsgChannel)
	container.Register(cont, func(cont *container.Container) (*UserMsgChannel, error) {
		return &requestChan, nil
	})
}

type User struct {
	usrMsgIn <-chan models.UserMsg
	db       Database
	ctx      context.Context
}

func (User) tableName() string {
	return "usr.users"
}

func UserDependencyConstructor(cont *container.Container) (*User, error) {
	db := New("postgres", cont.GetDB())

	if db == nil {
		return nil, errors.New("db is null")
	}

	userMsg, err := container.Resolve[UserMsgChannel](cont)
	if err != nil {
		return nil, err
	}

	dependency := &User{
		ctx:      cont.GetContext(),
		db:       *db,
		usrMsgIn: *userMsg,
	}

	go func() {
		dependency.MonitorChannels()
	}()

	return dependency, nil
}

func (repository User) MonitorChannels() {
	for {
		select {
		case request := <-repository.usrMsgIn:
			repository.replyRequest(request)
		case <-repository.ctx.Done():
			return
		}
	}
}

func (service User) replyRequest(req models.UserMsg) {
	switch {
	case req.Single != nil:
		requestBody := req.Single.RequestBody()
		if requestBody.ID != nil {
			user, err := service.GetByID(*requestBody.ID)
			if err != nil {
				req.Single.ReplyError(err)
				return
			}
			if user == nil {
				user = &models.User{}
			}
			req.Single.Reply(*user)
			return
		} else if requestBody.Email != nil {
			user, err := service.GetByEmail(*requestBody.Email)
			if err != nil {
				req.Single.ReplyError(err)
				return
			}
			if user == nil {
				user = &models.User{}
			}
			req.Single.Reply(*user)
			return
		}
	case req.List != nil:
		requestBody := req.List.RequestBody()
		userID := req.List.UserID()
		resp, err := service.Get(userID, requestBody)
		if err != nil {
			req.List.ReplyError(err)
			return
		}
		req.List.Reply(*resp)
	case req.Create != nil:
		requestBody := req.Create.RequestBody()

		userID := req.Create.UserID()

		err := service.Create(userID, requestBody)
		if err != nil {
			req.Create.ReplyError(err)
			return
		}
		req.Create.Reply(request.Success())
	case req.Update != nil:
		requestBody := req.Update.RequestBody()
		userID := req.Update.UserID()
		err := service.Update(userID, requestBody)
		if err != nil {
			req.Update.ReplyError(err)
			return
		}
		req.Update.Reply(request.Success())
	case req.Delete != nil:
		requestBody := req.Delete.RequestBody()
		userID := req.Delete.UserID()
		err := service.Delete(userID, requestBody)
		if err != nil {
			req.Delete.ReplyError(err)
			return
		}
		req.Delete.Reply(request.Success())
	}
}

func (repo User) Create(userID string, users []models.User) (err error) {
	count, err := sqlbuilder.GetInsertWithPermissions(repo.db, "USER", userID)
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("there is no effective permission to create this resource")
	}

	for i := range users {
		users[i].RegisteredAt = pointers.Time(time.Now())
	}
	_, err = repo.db.Insert(repo.tableName()).Rows(users).Executor().Exec()
	return
}

func (repo User) GetByID(id string) (*models.User, error) {
	user := models.User{}

	query := repo.db.From(repo.tableName()).Where(Ex{"id": id})

	_, err := query.ScanStruct(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (repo User) GetByEmail(email string) (*models.User, error) {
	user := models.User{}

	query := repo.db.From(repo.tableName()).Where(Ex{"email": email})

	_, err := query.ScanStruct(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (repo User) Get(userID string, request models.UserListRequestBody) (*models.UserListResponse, error) {
	data := []models.User{}

	query := repo.db.From(repo.tableName())

	if request.Filter.Search != nil {
		pattern := fmt.Sprintf("%%%s%%", *request.Filter.Search)
		query = query.Where(
			Or(
				I("first_name").Like(pattern),
				I("last_name").Like(pattern),
				I("email").Like(pattern),
			))
	}

	query = sqlbuilder.GetSelectWithPermissions(
		query,
		"USER",
		I("users.id"),
		userID,
	)

	count, err := query.Count()
	if err != nil {
		return nil, err
	}

	if request.Pagination.Page == nil {
		request.Pagination.Page = pointers.UInt(models.DefaultPage)
	}

	if request.Pagination.PageSize != nil {
		page := *request.Pagination.Page
		if page > 0 {
			page--
		}

		query = query.Offset(page * *request.Pagination.PageSize)
		query = query.Limit(*request.Pagination.PageSize)
	}

	err = query.ScanStructs(&data)
	if err != nil {
		return nil, err
	}

	result := models.UserListResponse{
		Count:    count,
		PageSize: request.Pagination.PageSize,
		Page:     request.Pagination.Page,
		Data:     data,
	}

	return &result, nil
}

func (repo User) Update(userID string, user models.User) error {
	// user.UpdatedAt = pointers.Time(time.Now())
	query := repo.db.Update(repo.tableName()).
		Set(user).
		Where(C("id").Eq(*user.ID))

	query = sqlbuilder.GetUpdateWithPermissions(
		query,
		"USER",
		I("users.id"),
		userID,
	)

	_, err := query.Executor().Exec()
	return err
}

func (repo User) Delete(userID, id string) error {
	query := repo.db.Delete(repo.tableName()).
		Where(C("id").Eq(id))

	query = sqlbuilder.GetDeleteWithPermissions(
		query,
		"USER",
		I("users.id"),
		userID,
	)

	result, err := query.Executor().Exec()
	if err != nil {
		return err
	}

	effectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if effectedRows == 0 {
		return errors.New("no rows have been effected")
	}

	return err
}
