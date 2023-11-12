package types

import (
	"strconv"
)

const nullInt = "NullInt"

type NullInt struct {
	BaseType[int]
}

func NewInt(value int, valid bool) NullInt {
	base := New[int](value, true)
	base.name = nullInt
	return NullInt{
		BaseType: base,
	}
}

func IntFrom(value int) NullInt {
	return NewInt(value, true)
}

func IntFromPtr(value *int) NullInt {
	base := FromPtr[int](value)
	base.name = nullInt
	return NullInt{
		BaseType: base,
	}
}

func (ttype NullInt) ImplementsGraphQLType(name string) bool {
	return nullInt == name
}

func (ttype *NullInt) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case int:
		ttype.Data = input
		return nil
	case int32:
		ttype.Data = int(input)
		return nil
	case int64:
		ttype.Data = int(input)
		return nil
	case string:
		value, err := strconv.Atoi(input)
		if err != nil {
			return err
		}
		ttype.Data = value
		ttype.Valid = true
		ttype.Set = true
		return nil
	default:
		return ttype.BaseType.UnmarshalGraphQL(input)
	}
}
