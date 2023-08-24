package request

import (
	"context"
	"errors"
	"time"
)

func Success() struct{} {
	return struct{}{}
}

type Signal struct{}

type RequestHandlerOption[RequestT any, RespT any] func(*RequestHandler[RequestT, RespT])

type RequestHandler[RequestParamsT any, RespT any] struct {
	ctx           context.Context
	userID        string
	response      response[RespT]
	requestParams RequestParamsT

	timeoutInterval         time.Duration
	responseTimeoutInterval time.Duration
}

func New[RequestT any, RespT any](ctx context.Context, body RequestT, opts ...RequestHandlerOption[RequestT, RespT]) RequestHandler[RequestT, RespT] {
	req := RequestHandler[RequestT, RespT]{
		ctx:                     ctx,
		response:                newResponse[RespT](),
		timeoutInterval:         DefaultTimeOut,
		responseTimeoutInterval: DefaultTimeOut,
		requestParams:           body,
	}

	for _, opt := range opts {
		opt(&req)
	}

	return req
}

func (r RequestHandler[RequestT, RespT]) Context() context.Context {
	return r.ctx
}

func (r RequestHandler[RequestT, RespT]) UserID() string {
	return r.userID
}

func (r RequestHandler[RequestT, RespT]) RequestParams() RequestT {
	return r.requestParams
}

func (r RequestHandler[RequestT, RespT]) Wait() (*RespT, error) {
	ticker := time.NewTicker(r.timeoutInterval)
	defer ticker.Stop()

	select {
	case msg := <-r.response.result:
		return &msg, nil
	case err := <-r.response.err:
		return nil, err
	case <-ticker.C:
		return nil, errors.New(requestTimedout)
	case <-r.ctx.Done():
		return nil, errors.New(contextCancelled)
	}
}

func (r *RequestHandler[RequestT, RespT]) SetOptions(opts ...RequestHandlerOption[RequestT, RespT]) {
	for _, opt := range opts {
		opt(r)
	}
}

func (r RequestHandler[RequestT, RespT]) Reply(response RespT) {
	ticker := time.NewTicker(r.timeoutInterval)
	defer ticker.Stop()

	go func() {
		select {
		case r.response.result <- response:
		case <-ticker.C:
		}
	}()
}

func (r RequestHandler[RequestT, RespT]) ReplyError(err error) {
	ticker := time.NewTicker(r.timeoutInterval)
	defer ticker.Stop()

	go func() {
		select {
		case r.response.err <- err:
		case <-ticker.C:
		}
	}()
}

func TimeoutOption[RequestT any, RespT any](timeout time.Duration) RequestHandlerOption[RequestT, RespT] {
	return func(r *RequestHandler[RequestT, RespT]) {
		r.timeoutInterval = timeout
	}
}

func UserIDOption[RequestT any, RespT any](userID string) RequestHandlerOption[RequestT, RespT] {
	return func(r *RequestHandler[RequestT, RespT]) {
		r.userID = userID
	}
}

type response[T any] struct {
	err    chan error `json:"-"`
	result chan T     `json:"-"`
}

func newResponse[T any]() response[T] {
	err := make(chan error)
	result := make(chan T)
	return response[T]{
		err:    err,
		result: result,
	}
}
