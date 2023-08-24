package request

import (
	"context"
	"errors"
	"time"
)

func CreateSenderFunc[T any](ch chan<- T, timeoutInterval time.Duration) func(ctx context.Context, msg T) error {
	return func(ctx context.Context, msg T) error {
		ticker := time.NewTicker(timeoutInterval)
		defer ticker.Stop()
		select {
		case ch <- msg:
			return nil
		case <-ticker.C:
			return errors.New(requestSendTimedout)
		case <-ctx.Done():
			return errors.New(contextCancelled)
		}
	}
}
