package types

import (
	"encoding/json"
	"time"
)

const timeName = "Time"

type Time struct {
	BaseType[time.Time]
}

func NewTime(value time.Time, valid bool) Time {
	base := New[time.Time](value, true)
	base.name = timeName
	return Time{
		BaseType: base,
	}
}

func TimeFrom(value time.Time) Time {
	return NewTime(value, true)
}

func TimeNow() Time {
	return TimeFrom(time.Now())
}

func TimeFromPtr(value *time.Time) Time {
	base := FromPtr[time.Time](value)
	base.name = timeName
	return Time{
		BaseType: base,
	}
}

func (ttype Time) ImplementsGraphQLType(name string) bool {
	return timeName == name
}

func (ttype *Time) UnmarshalGraphQL(input interface{}) error {
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

func (ttype Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(ttype.Data)
}
