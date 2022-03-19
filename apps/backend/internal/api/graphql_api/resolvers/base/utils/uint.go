package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Uint struct {
	uint
}

func NewUint(v uint) Uint {
	return Uint{v}
}

func (u Uint) Value() uint {
	return u.uint
}

// ImplementsGraphQLType maps this custom Go type
// to the graphql scalar type in the schema.
func (Uint) ImplementsGraphQLType(name string) bool {
	return name == "Uint"
}

// UnmarshalGraphQL is a custom unmarshaler for Int
//
// This function will be called whenever you use the
// Uint scalar as an input
func (i *Uint) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case uint:
		i.uint = input
		return nil
	case int32:
		i.uint = uint(input)
		return nil
	case int64:
		i.uint = uint(input)
		return nil
	case string:
		int, err := strconv.Atoi(input)
		if err != nil {
			return err
		}
		i.uint = uint(int)
		return nil
	default:
		return fmt.Errorf("wrong type for Uint: %T", input)
	}
}

// MarshalJSON is a custom marshaler for Uint
//
// This function will be called whenever you
// query for fields that use the Uint type
func (i Uint) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.uint)
}
