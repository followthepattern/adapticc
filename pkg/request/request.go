package request

import (
	"context"
	"errors"
	"time"
)

var Success = struct{}{}

type Signal struct{}

type TaskOption[TaskT any, SuccessResultT any] func(*Task[TaskT, SuccessResultT])

type Task[RequestParamsT any, RespT any] struct {
	ctx           context.Context
	userID        string
	result        result[RespT]
	requestParams RequestParamsT

	timeoutInterval         time.Duration
	responseTimeoutInterval time.Duration
}

func New[TaskT any, SuccessResultT any](ctx context.Context, body TaskT, opts ...TaskOption[TaskT, SuccessResultT]) Task[TaskT, SuccessResultT] {
	req := Task[TaskT, SuccessResultT]{
		ctx:                     ctx,
		result:                  newResult[SuccessResultT](),
		timeoutInterval:         DefaultTimeOut,
		responseTimeoutInterval: DefaultTimeOut,
		requestParams:           body,
	}

	for _, opt := range opts {
		opt(&req)
	}

	return req
}

func (r Task[TaskT, SuccessResultT]) Context() context.Context {
	return r.ctx
}

func (r Task[TaskT, SuccessResultT]) UserID() string {
	return r.userID
}

func (r Task[TaskT, SuccessResultT]) RequestParams() TaskT {
	return r.requestParams
}

func (r Task[TaskT, SuccessResultT]) Wait() (*SuccessResultT, error) {
	ticker := time.NewTicker(r.timeoutInterval)
	defer ticker.Stop()

	successResult := r.result.successResult
	errResult := r.result.err

	select {
	case msg := <-successResult:
		return &msg, nil
	case err := <-errResult:
		return nil, err
	case <-ticker.C:
		return nil, errors.New(requestTimedout)
	case <-r.ctx.Done():
		return nil, errors.New(contextCancelled)
	}
}

func (r *Task[TaskT, SuccessResultT]) SetOptions(opts ...TaskOption[TaskT, SuccessResultT]) {
	for _, opt := range opts {
		opt(r)
	}
}

func (r Task[TaskT, SuccessResultT]) Reply(response SuccessResultT) {
	ticker := time.NewTicker(r.timeoutInterval)
	defer ticker.Stop()

	successResult := r.result.successResult

	go func() {
		select {
		case successResult <- response:
		case <-ticker.C:
		}
	}()
}

func (r Task[TaskT, SuccessResultT]) ReplyError(err error) {
	ticker := time.NewTicker(r.timeoutInterval)
	defer ticker.Stop()

	errResult := r.result.err

	go func() {
		select {
		case errResult <- err:
		case <-ticker.C:
		}
	}()
}

func TimeoutOption[TaskT any, SuccessResultT any](timeout time.Duration) TaskOption[TaskT, SuccessResultT] {
	return func(r *Task[TaskT, SuccessResultT]) {
		r.timeoutInterval = timeout
	}
}

func UserIDOption[TaskT any, SuccessResultT any](userID string) TaskOption[TaskT, SuccessResultT] {
	return func(r *Task[TaskT, SuccessResultT]) {
		r.userID = userID
	}
}

type result[T any] struct {
	err           chan error `json:"-"`
	successResult chan T     `json:"-"`
}

func newResult[T any]() result[T] {
	err := make(chan error)
	successResult := make(chan T)
	return result[T]{
		err:           err,
		successResult: successResult,
	}
}
