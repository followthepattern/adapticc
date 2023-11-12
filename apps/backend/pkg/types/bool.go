package types

import (
	"strconv"
)

const nullBool = "NullBool"

type NullBool struct {
	BaseType[bool]
}

func NewBool(value bool, valid bool) NullBool {
	base := New[bool](value, true)
	base.name = nullBool
	return NullBool{
		BaseType: base,
	}
}

func BoolFrom(value bool) NullBool {
	return NewBool(value, true)
}

func BoolFromPtr(value *bool) NullBool {
	base := FromPtr[bool](value)
	base.name = nullBool
	return NullBool{
		BaseType: base,
	}
}

func (ttype NullBool) ImplementsGraphQLType(name string) bool {
	return nullBool == name
}

func (ttype *NullBool) UnmarshalGraphQL(input interface{}) error {
	ttype.name = nullBool
	switch input := input.(type) {
	case string:
		value, err := strconv.ParseBool(input)
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
