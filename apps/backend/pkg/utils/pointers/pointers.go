package pointers

import "time"

func ToPtr[T any](v T) *T {
	return &v
}

func String(s string) *string {
	return &s
}

func Time(t time.Time) *time.Time {
	return &t
}

func Bool(b bool) *bool {
	return &b
}

func Int(i int) *int {
	return &i
}

func UInt(ui uint) *uint {
	return &ui
}