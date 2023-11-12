package types

import "strconv"

const nullInt64 = "NullInt64"

type NullInt64 struct {
	BaseType[int64]
}

func NewInt64(value int64, valid bool) NullInt64 {
	base := New[int64](value, true)
	base.name = nullInt64
	return NullInt64{
		BaseType: base,
	}
}

func Int64From(value int64) NullInt64 {
	return NewInt64(value, true)
}

func Int64FromPtr(value *int64) NullInt64 {
	base := FromPtr[int64](value)
	base.name = nullInt64
	return NullInt64{
		BaseType: base,
	}
}

func (ttype NullInt64) ImplementsGraphQLType(name string) bool {
	return nullInt64 == name
}

func (ttype *NullInt64) UnmarshalGraphQL(input interface{}) error {
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
