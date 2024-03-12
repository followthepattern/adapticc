package task

import (
	"context"
	"errors"
	"time"
)

func Success() struct{} {
	return struct{}{}
}

const (
	contextCancelled = "context is canceled"
	taskTimedOut     = "task timed out"
	taskRunTimedOut  = "task sending timed out"

	DefaultTimeOut time.Duration = time.Second * 3
)

type Signal struct{}

type TaskOption[ResultT any] func(*Task[ResultT])

type TaskAction[ResultT any] func() (ResultT, error)

type Task[ResultT any] struct {
	ctx    context.Context
	action TaskAction[ResultT]

	err    chan error
	result chan ResultT

	timeoutInterval         time.Duration
	responseTimeoutInterval time.Duration
}

func New[ResultT any](action TaskAction[ResultT], opts ...TaskOption[ResultT]) Task[ResultT] {
	req := Task[ResultT]{
		ctx:                     context.Background(),
		timeoutInterval:         DefaultTimeOut,
		responseTimeoutInterval: DefaultTimeOut,
		action:                  action,
		err:                     make(chan error),
		result:                  make(chan ResultT),
	}

	for _, opt := range opts {
		opt(&req)
	}

	return req
}

func (t Task[ResultT]) reply(result ResultT) {
	ticker := time.NewTicker(t.responseTimeoutInterval)
	defer ticker.Stop()

	select {
	case t.result <- result:
	case <-t.ctx.Done():
	case <-ticker.C:
	}

}

func (t Task[ResultT]) Run() Task[ResultT] {
	go func() {
		res, err := t.action()
		if err != nil {
			t.replyError(err)
			return
		}

		t.reply(res)
	}()

	return t
}

func (t Task[RequestT]) replyError(err error) {
	ticker := time.NewTicker(t.timeoutInterval)
	defer ticker.Stop()

	select {
	case t.err <- err:
	case <-t.ctx.Done():
	case <-ticker.C:
	}

}

func (r Task[ResultT]) Wait() (*ResultT, error) {
	ticker := time.NewTicker(r.timeoutInterval)
	defer ticker.Stop()

	select {
	case msg := <-r.result:
		return &msg, nil
	case err := <-r.err:
		return nil, err
	case <-ticker.C:
		return nil, errors.New(taskTimedOut)
	case <-r.ctx.Done():
		return nil, errors.New(contextCancelled)
	}
}

func TimeoutOption[ResultT any, RespT any](timeout time.Duration) TaskOption[ResultT] {
	return func(r *Task[ResultT]) {
		r.timeoutInterval = timeout
	}
}

func ContextOption[ResultT any, RespT any](ctx context.Context) TaskOption[ResultT] {
	return func(r *Task[ResultT]) {
		r.ctx = ctx
	}
}
