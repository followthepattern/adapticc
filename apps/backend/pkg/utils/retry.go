package utils

import (
	"context"
	"time"
)

const (
	incrementQuartile       = 4
	resultIncrementInterval = 30 * time.Minute
)

func Retry(ctx context.Context, interval, until time.Duration, call func() bool) bool {
	var counter int64
	untilTimer := time.NewTimer(until)

	counter++
	if ok := call(); ok {
		return true
	}
	intervalTimer := time.NewTimer(incrementInterval(interval, counter))

	for {
		select {
		case <-untilTimer.C:
			return false
		case <-intervalTimer.C:
			counter++
			if ok := call(); ok {
				return true
			}
			intervalTimer.Reset(incrementInterval(interval, counter))
		case <-ctx.Done():
			return false
		}
	}
}

func incrementInterval(interval time.Duration, counter int64) time.Duration {
	quartile := int64(interval) / incrementQuartile
	addendum := quartile * counter
	result := interval + time.Duration(addendum)

	if result > resultIncrementInterval {
		result = resultIncrementInterval
	}
	return result
}
