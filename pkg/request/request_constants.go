package request

import "time"

const (
	contextCancelled    = "context is canceled"
	requestTimedout     = "request timed out"
	requestSendTimedout = "request sending timed out"

	DefaultTimeOut time.Duration = time.Second * 3
)
