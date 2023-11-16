package types

import (
	"strconv"
)

const uinName = "Uint"

type Uint struct {
	BaseType[uint]
}

func NewUint(value uint, valid bool) Uint {
	base := New[uint](value, true)
	base.name = uinName
	return Uint{
		BaseType: base,
	}
}

func UintFrom(value uint) Uint {
	return NewUint(value, true)
}

func UintFromPtr(value *uint) Uint {
	base := FromPtr[uint](value)
	base.name = uinName
	return Uint{
		BaseType: base,
	}
}

func (ttype Uint) ImplementsGraphQLType(name string) bool {
	return name == uinName
}

func (ttype *Uint) UnmarshalGraphQL(input interface{}) error {
	ttype.name = uinName
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
