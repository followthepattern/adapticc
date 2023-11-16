package types

import (
	"strconv"
)

const nullUint = "NullUint"

type NullUint struct {
	BaseType[uint]
}

func NewUint(value uint, valid bool) NullUint {
	base := New[uint](value, true)
	base.name = nullUint
	return NullUint{
		BaseType: base,
	}
}

func UintFrom(value uint) NullUint {
	return NewUint(value, true)
}

func UintFromPtr(value *uint) NullUint {
	base := FromPtr[uint](value)
	base.name = nullUint
	return NullUint{
		BaseType: base,
	}
}

func (ttype NullUint) ImplementsGraphQLType(name string) bool {
	return name == nullUint
}

func (ttype *NullUint) UnmarshalGraphQL(input interface{}) error {
	ttype.name = nullUint
	switch input := input.(type) {
	case int:
		ttype.Data = uint(input)
		ttype.Valid = true
		ttype.Set = true
		return nil
	case int32:
		ttype.Data = uint(input)
		ttype.Valid = true
		ttype.Set = true
		return nil
	case int64:
		ttype.Data = uint(input)
		ttype.Valid = true
		ttype.Set = true
		return nil
	case string:
		value, err := strconv.Atoi(input)
		if err != nil {
			return err
		}
		ttype.Data = uint(value)
		ttype.Valid = true
		ttype.Set = true
		return nil
	default:
		return ttype.BaseType.UnmarshalGraphQL(input)
	}
}
