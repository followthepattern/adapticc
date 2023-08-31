package request

import "time"

const (
	contextCancelled = "context is canceled"
	taskTimedout     = "task timed out"
	taskSendTimedout = "task sending timed out"

	DefaultTimeOut time.Duration = time.Second * 3
)
