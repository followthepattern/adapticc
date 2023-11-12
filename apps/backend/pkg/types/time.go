package types

import (
	"time"
)

const nullTime = "NullTime"

type NullTime struct {
	BaseType[time.Time]
}

func NewTime(value time.Time, valid bool) NullTime {
	base := New[time.Time](value, true)
	base.name = nullTime
	return NullTime{
		BaseType: base,
	}
}

func TimeFrom(value time.Time) NullTime {
	return NewTime(value, true)
}

func TimeFromPtr(value *time.Time) NullTime {
	base := FromPtr[time.Time](value)
	base.name = nullTime
	return NullTime{
		BaseType: base,
	}
}

func (ttype NullTime) ImplementsGraphQLType(name string) bool {
	return nullTime == name
}

func (ttype *NullTime) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case time.Time:
		ttype.Data = input
		ttype.Valid = true
		ttype.Set = true
		return nil
	case string:
		data, err := time.Parse(time.RFC3339, input)
		if err != nil {
			return err
		}
		ttype.Data = data
		ttype.Valid = true
		ttype.Set = true
		return nil
	case []byte:
		data, err := time.Parse(time.RFC3339, string(input))
		if err != nil {
			return err
		}
		ttype.Data = data
		ttype.Valid = true
		ttype.Set = true
		return nil
	case int32:
		ttype.Data = time.Unix(int64(input), 0)
		ttype.Valid = true
		ttype.Set = true
		return nil
	case int64:
		if input >= 1e10 {
			sec := input / 1e9
			nsec := input - (sec * 1e9)
			ttype.Data = time.Unix(sec, nsec)
		} else {
			ttype.Data = time.Unix(input, 0)
		}
		ttype.Valid = true
		ttype.Set = true
		return nil
	case float64:
		ttype.Data = time.Unix(int64(input), 0)
		ttype.Valid = true
		ttype.Set = true
		return nil
	default:
		return ttype.BaseType.UnmarshalGraphQL(input)
	}
}
