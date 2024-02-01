package types

import "strconv"

const int64Name = "Int64"

type Int64 struct {
	BaseType[int64]
}

func NewInt64(value int64, valid bool) Int64 {
	base := New[int64](value, true)
	base.name = int64Name
	return Int64{
		BaseType: base,
	}
}

func Int64From(value int64) Int64 {
	return NewInt64(value, true)
}

func Int64FromPtr(value *int64) Int64 {
	base := FromPtr[int64](value)
	base.name = int64Name
	return Int64{
		BaseType: base,
	}
}

func (ttype Int64) ImplementsGraphQLType(name string) bool {
	return int64Name == name
}

func (ttype *Int64) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case int:
		ttype.Data = int64(input)
		ttype.Valid = true
		ttype.Set = true
		return nil
	case int32:
		ttype.Data = int64(input)
		ttype.Valid = true
		ttype.Set = true
		return nil
	case int64:
		ttype.Data = input
		ttype.Valid = true
		ttype.Set = true
		return nil
	case string:
		value, err := strconv.Atoi(input)
		if err != nil {
			return err
		}
		ttype.Data = int64(value)
		ttype.Valid = true
		ttype.Set = true
		return nil
	default:
		return ttype.BaseType.UnmarshalGraphQL(input)
	}
}
