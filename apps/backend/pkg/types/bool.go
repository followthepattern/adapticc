package types

import (
	"strconv"
)

const boolName = "Boolean"

var (
	TRUE  = BoolFrom(true)
	FALSE = BoolFrom(false)
)

type Bool struct {
	BaseType[bool]
}

func NewBool(value bool, valid bool) Bool {
	base := New[bool](value, true)
	base.name = boolName
	return Bool{
		BaseType: base,
	}
}

func BoolFrom(value bool) Bool {
	return NewBool(value, true)
}

func BoolFromPtr(value *bool) Bool {
	base := FromPtr[bool](value)
	base.name = boolName
	return Bool{
		BaseType: base,
	}
}

func (ttype Bool) ImplementsGraphQLType(name string) bool {
	return boolName == name
}

func (ttype *Bool) UnmarshalGraphQL(input interface{}) error {
	ttype.name = boolName
	switch input := input.(type) {
	case bool:
		ttype.Data = input
		ttype.Valid = true
		ttype.Set = true
		return nil
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
