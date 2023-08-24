package services

import (
	"context"

	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/followthepattern/adapticc/pkg/models"
	repositories "github.com/followthepattern/adapticc/pkg/repositories/database"
	"github.com/followthepattern/adapticc/pkg/request"
)

type UserMsgChannel chan models.UserMsg

func RegisterUserChannel(cont *container.Container) {
	if cont == nil {
		return
	}
	userMsgChannel := make(UserMsgChannel)
	container.Register(cont, func(cont *container.Container) (*UserMsgChannel, error) {
		return &userMsgChannel, nil
	})
}

type User struct {
	userMsgChannelIn  <-chan models.UserMsg
	userMsgChannelOut chan<- models.UserMsg
	ctx               context.Context
	sendUserMsg       func(ctx context.Context, msg models.UserMsg) error
}

func UserDependencyConstructor(cont *container.Container) (*User, error) {
	userMsgChannelIn, err := container.Resolve[UserMsgChannel](cont)
	if err != nil {
		return nil, err
	}

	userMsgChannelOut, err := container.Resolve[repositories.UserMsgChannel](cont)
	if err != nil {
		return nil, err
	}

	dependency := User{
		ctx:               cont.GetContext(),
		userMsgChannelIn:  *userMsgChannelIn,
		userMsgChannelOut: *userMsgChannelOut,
		sendUserMsg:       request.CreateSenderFunc(*userMsgChannelOut, request.DefaultTimeOut),
	}

	go func() {
		dependency.MonitorChannels()

	}()

	return &dependency, nil
}

func (service User) MonitorChannels() {
	for {
		select {
		case request := <-service.userMsgChannelIn:
			service.replyUserMsgRequest(request)
		case <-service.ctx.Done():
			return
		}
	}
}

func (service User) replyUserMsgRequest(request models.UserMsg) {
	service.sendUserMsg(service.ctx, request)
}
