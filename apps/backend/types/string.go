package types

import (
	"encoding/json"
	"fmt"
)

const stringName = "String"

type String struct {
	BaseType[string]
}

func NewString(value string, valid bool) String {
	base := New[string](value, true)
	base.name = stringName
	return String{
		BaseType: base,
	}
}

func StringFrom(value string) String {
	return NewString(value, true)
}

func StringFromPtr(value *string) String {
	base := FromPtr[string](value)
	base.name = stringName
	return String{
		BaseType: base,
	}
}

func (ttype String) ImplementsGraphQLType(name string) bool {
	return stringName == name
}

func (ttype *String) UnmarshalGraphQL(input interface{}) error {
	ttype.name = stringName
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

func (ttype String) Len() int {
	if ttype.IsValid() {
		return len(ttype.Data)
	}
	return 0
}

func (ttype String) MarshalJSON() ([]byte, error) {
	return json.Marshal(ttype.Data)
}
