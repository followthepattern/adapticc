package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/followthepattern/adapticc/pkg/models"
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
		resp, err := service.Get(requestBody)
		if err != nil {
			req.List.ReplyError(err)
			return
		}
		req.List.Reply(*resp)
	case req.Create != nil:
		requestBody := req.Create.RequestBody()
		err := service.Create(requestBody)
		if err != nil {
			req.Create.ReplyError(err)
			return
		}
		req.Create.Reply(request.Success())
	case req.Update != nil:
		requestBody := req.Update.RequestBody()
		err := service.Update(requestBody)
		if err != nil {
			req.Update.ReplyError(err)
			return
		}
		req.Update.Reply(request.Success())
	case req.Delete != nil:
		requestBody := req.Delete.RequestBody()
		err := service.Delete(requestBody)
		if err != nil {
			req.Delete.ReplyError(err)
			return
		}
		req.Delete.Reply(request.Success())
	}
}

func (repo User) Create(users []models.User) (err error) {
	// for _, user := range users {
	// 	user.CreatedAt = pointers.Time(time.Now())
	// }
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

func (repo User) Get(request models.UserListRequestBody) (*models.UserListResponse, error) {
	users := []models.User{}

	query := repo.db.From(repo.tableName())

	if request.Search != nil {
		pattern := fmt.Sprintf("%%%v%%", *request.Search)
		query = query.Where(
			Or(
				I("first_name").Like(pattern),
				I("last_name").Like(pattern),
				I("email").Like(pattern),
			))
	}

	count, err := query.Count()
	if err != nil {
		return nil, err
	}

	result := models.UserListResponse{
		Count:    count,
		PageSize: request.PageSize,
	}

	if request.Page == nil {
		result.Page = pointers.UInt(1)
	}

	if request.PageSize != nil {
		page := *request.Page
		if page > 0 {
			page--
		}

		query = query.Offset(page * *request.PageSize)
		query = query.Limit(*request.PageSize)
	}

	err = query.ScanStructs(&users)
	if err != nil {
		return nil, err
	}

	result.Data = users

	return &result, nil
}

func (repo User) Update(user models.User) error {
	// user.UpdatedAt = pointers.Time(time.Now())
	_, err := repo.db.Update(repo.tableName()).
		Set(user).
		Where(C("id").Eq(*user.ID)).
		Executor().Exec()
	return err
}

func (repo User) Delete(id string) error {
	_, err := repo.db.Delete(repo.tableName()).Executor().Exec()
	return err
}
