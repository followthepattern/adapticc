package resolvers

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Int64 struct {
	int64
}

func NewInt64(v int64) Int64 {
	return Int64{v}
}

func (i Int64) Value() int64 {
	return i.int64
}

// ImplementsGraphQLType maps this custom Go type
// to the graphql scalar type in the schema.
func (Int64) ImplementsGraphQLType(name string) bool {
	return name == "Int64"
}

// UnmarshalGraphQL is a custom unmarshaler for Int
//
// This function will be called whenever you use the
// Int64 scalar as an input
func (i *Int64) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case int32:
		i.int64 = int64(input)
		return nil
	case int64:
		i.int64 = input
		return nil
	case string:
		int, err := strconv.Atoi(input)
		if err != nil {
			return err
		}
		i.int64 = int64(int)
		return nil
	default:
		return fmt.Errorf("wrong type for Uint: %T", input)
	}
}

// MarshalJSON is a custom marshaler for Int64
//
// This function will be called whenever you
// query for fields that use the Int64 type
func (i Int64) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.int64)
}

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
	case float64:
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

func (i *Uint) UnmarshalJSON(data []byte) error {
	var v uint
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	i.uint = v
	return nil
}
