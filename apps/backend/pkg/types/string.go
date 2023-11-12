package types

import (
	"fmt"
)

const nullString = "NullString"

type NullString struct {
	BaseType[string]
}

func NewString(value string, valid bool) NullString {
	base := New[string](value, true)
	base.name = nullString
	return NullString{
		BaseType: base,
	}
}

func StringFrom(value string) NullString {
	return NewString(value, true)
}

func StringFromPtr(value *string) NullString {
	base := FromPtr[string](value)
	base.name = nullString
	return NullString{
		BaseType: base,
	}
}

func (ttype NullString) ImplementsGraphQLType(name string) bool {
	return nullString == name
}

func (ttype *NullString) UnmarshalGraphQL(input interface{}) error {
	ttype.name = nullString
	switch input := input.(type) {
	case string:
		ttype.Data = fmt.Sprintf("%v", input)
		ttype.Valid = true
		ttype.Set = true
		return nil
	default:
		return ttype.BaseType.UnmarshalGraphQL(input)
	}
}

func (ttype NullString) Len() int {
	if ttype.IsValid() {
		return len(ttype.Data)
	}
	return 0
}
