package services

import (
	"context"

	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/followthepattern/adapticc/pkg/models"
	repositories "github.com/followthepattern/adapticc/pkg/repositories/database"
	"github.com/followthepattern/adapticc/pkg/request"
)

type Product struct {
	productMsgChannelIn  <-chan models.ProductMsg
	productMsgChannelOut chan<- models.ProductMsg
	ctx                  context.Context
	cfg                  config.Config
	sendProductMsg       func(ctx context.Context, msg models.ProductMsg) error
}

type ProductMsgChannel chan models.ProductMsg

func RegisterProductChannel(cont *container.Container) {
	if cont == nil {
		return
	}
	productMsgChannel := make(ProductMsgChannel)
	container.Register(cont, func(cont *container.Container) (*ProductMsgChannel, error) {
		return &productMsgChannel, nil
	})
}

func ProductDependencyConstructor(cont *container.Container) (*Product, error) {
	productMsgChannelIn, err := container.Resolve[ProductMsgChannel](cont)
	if err != nil {
		return nil, err
	}

	productMsgChannelOut, err := container.Resolve[repositories.ProductMsgChannel](cont)
	if err != nil {
		return nil, err
	}

	dependency := Product{
		ctx:                  cont.GetContext(),
		cfg:                  cont.GetConfig(),
		productMsgChannelIn:  *productMsgChannelIn,
		productMsgChannelOut: *productMsgChannelOut,
		sendProductMsg:       request.CreateSenderFunc(*productMsgChannelOut, request.DefaultTimeOut),
	}

	go func() {
		dependency.MonitorChannels()
	}()

	return &dependency, nil
}

func (service Product) MonitorChannels() {
	for {
		select {
		case msg := <-service.productMsgChannelIn:
			switch {
			case msg.Single != nil:
				service.sendProductMsg(msg.Single.Context(), msg)
			case msg.List != nil:
				service.sendProductMsg(msg.List.Context(), msg)
			case msg.Create != nil:
				service.sendProductMsg(msg.Create.Context(), msg)
			case msg.Update != nil:
				service.sendProductMsg(msg.Update.Context(), msg)
			case msg.Delete != nil:
				service.sendProductMsg(msg.Delete.Context(), msg)
			}
		case <-service.ctx.Done():
			return
		}
	}
}
