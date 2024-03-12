package types

import (
	"strconv"
)

const intName = "Int"

type Int struct {
	BaseType[int]
}

func NewInt(value int, valid bool) Int {
	base := New[int](value, true)
	base.name = intName
	return Int{
		BaseType: base,
	}
}

func IntFrom(value int) Int {
	return NewInt(value, true)
}

func IntFromPtr(value *int) Int {
	base := FromPtr[int](value)
	base.name = intName
	return Int{
		BaseType: base,
	}
}

func (ttype Int) ImplementsGraphQLType(name string) bool {
	return intName == name
}

func (ttype *Int) UnmarshalGraphQL(input interface{}) error {
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
