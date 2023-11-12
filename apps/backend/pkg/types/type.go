package types

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type BaseType[T any] struct {
	Data  T
	Valid bool
	Set   bool
	name  string
}

func (*BaseType[T]) Nullable() {}

func (ttype BaseType[T]) UnmarshalGraphQL(input interface{}) error {
	return fmt.Errorf("wrong type for %s: %T", ttype.name, input)
}

func FromType[T any](value T) BaseType[T] {
	return New(value, true)
}

func FromPtr[T any](s *T) BaseType[T] {
	if s == nil {
		var value T
		return New(value, false)
	}
	return New(*s, true)
}

func New[T any](value T, valid bool) BaseType[T] {
	return BaseType[T]{
		Data:  value,
		Valid: valid,
		Set:   true,
	}
}

func (ttype BaseType[T]) IsValid() bool {
	return ttype.Set && ttype.Valid
}

func (ttype BaseType[T]) IsSet() bool {
	return ttype.Set
}

func (ttype *BaseType[T]) UnmarshalJSON(Value []byte) error {
	ttype.Set = true
	if bytes.Equal(Value, NullBytes) {
		var value T
		ttype.Data = value
		ttype.Valid = false
		return nil
	}

	if err := json.Unmarshal(Value, &ttype.Data); err != nil {
		return err
	}

	ttype.Valid = true
	return nil
}

func (ttype BaseType[T]) MarshalJSON() ([]byte, error) {
	if !ttype.Valid {
		return NullBytes, nil
	}
	return json.Marshal(ttype.Data)
}

func (ttype BaseType[T]) MarshalText() ([]byte, error) {
	if !ttype.Valid {
		return NullBytes, nil
	}
	return []byte(fmt.Sprintf("%v", ttype.Data)), nil
}

func (ttype BaseType[T]) String() string {
	return fmt.Sprintf("%v", ttype.Data)
}

func (ttype *BaseType[T]) SetValid(value T) {
	ttype.Data = value
	ttype.Valid = true
	ttype.Set = true
}

func (ttype BaseType[T]) Ptr() *T {
	if !ttype.Valid {
		return nil
	}
	return &ttype.Data
}

func (ttype BaseType[T]) IsZero() bool {
	return !ttype.Valid
}

func (ttype *BaseType[T]) Scan(value interface{}) error {
	if value == nil {
		var defaultValue T
		ttype.Data, ttype.Valid, ttype.Set = defaultValue, false, false
		return nil
	}
	ttype.Valid, ttype.Set = true, true
	return ConvertAssign(&ttype.Data, value)
}

func (ttype BaseType[T]) Value() (driver.Value, error) {
	if !ttype.Valid {
		return nil, nil
	}
	return ttype.Data, nil
}
