package database

import (
	"context"
	"errors"
	"time"

	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/request"
	"github.com/followthepattern/adapticc/pkg/utils/pointers"
	"go.uber.org/zap"

	. "github.com/doug-martin/goqu/v9"
)

type AuthMsgChannel chan models.AuthMsg

func RegisterAuthChannel(cont *container.Container) {
	if cont == nil {
		return
	}
	requestChan := make(AuthMsgChannel)
	container.Register(cont, func(cont *container.Container) (*AuthMsgChannel, error) {
		return &requestChan, nil
	})

}

type Auth struct {
	authMsgIn <-chan models.AuthMsg
	logger    zap.Logger
	db        Database
	ctx       context.Context
}

func AuthDependencyConstructor(cont *container.Container) (*Auth, error) {
	db := New("postgres", cont.GetDB())

	if db == nil {
		return nil, errors.New("db is null")
	}

	AuthMsg, err := container.Resolve[AuthMsgChannel](cont)
	if err != nil {
		return nil, err
	}

	dependency := &Auth{
		ctx:       cont.GetContext(),
		db:        *db,
		authMsgIn: *AuthMsg,
		logger:    *cont.GetLogger(),
	}

	go func() {
		dependency.MonitorChannels()
	}()

	return dependency, nil
}

func (service Auth) MonitorChannels() {
	for {
		select {
		case request := <-service.authMsgIn:
			service.replyRequest(request)
		case <-service.ctx.Done():
			return
		}
	}
}

func (service Auth) replyRequest(req models.AuthMsg) {
	switch {
	case req.VerifyEmail != nil:
		service.replyVerifyEmail(*req.VerifyEmail)
	case req.RegisterUser != nil:
		service.replyRegisterUser(*req.RegisterUser)
	case req.VerifyLogin != nil:
		service.replyVerifyLogin(*req.VerifyLogin)
	}
}

func (service Auth) replyVerifyEmail(handler request.RequestHandler[string, bool]) {
	email := handler.RequestParams()

	count, err := service.db.From("usr.users").Where(Ex{"email": email}).Count()
	if err != nil {
		service.logger.Error(err.Error())
		handler.ReplyError(err)
		return
	}

	handler.Reply(count == 0)
}

func (service Auth) replyRegisterUser(handler request.RequestHandler[models.AuthUser, request.Signal]) {
	registerUser := handler.RequestParams()

	registerUser.Userlog = models.Userlog{
		CreatedAt: pointers.ToPtr(time.Now()),
	}

	_, err := service.db.Insert("usr.users").Rows(registerUser).Executor().Exec()
	if err != nil {
		service.logger.Error(err.Error())
		handler.ReplyError(err)
		return
	}

	handler.Reply(request.Success)
}

func (service Auth) replyVerifyLogin(handler request.RequestHandler[string, models.AuthUser]) {
	authUser := models.AuthUser{}

	email := handler.RequestParams()

	query := service.db.From("usr.users").Where(Ex{"email": email})

	_, err := query.ScanStruct(&authUser)
	if err != nil {
		service.logger.Error(err.Error())
		handler.ReplyError(err)
		return
	}

	handler.Reply(authUser)
}
