package sqlexpectations

import (
	"database/sql/driver"
	"reflect"
	"strings"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/followthepattern/adapticc/pkg/utils/pointers"
)

func ModelToSQLMockRows[T any](value T) *sqlmock.Rows {
	tvalue := reflect.TypeOf(value)
	sqlItemType := getSqlItemType(value)
	rowCount := getSqlRowCount(value)

	columns := []string{}

	for i := 0; i < sqlItemType.NumField(); i++ {
		tag := strings.Split(sqlItemType.Field(i).Tag.Get("json"), ",")[0]
		if tag == "" || tag == "-" {
			tag = strings.Split(sqlItemType.Field(i).Tag.Get("db"), ",")[0]
			if tag == "" {
				continue
			}
		}

		if !IsSQLType(sqlItemType.Field(i).Type) {
			continue
		}
		columns = append(columns, tag)
	}

	result := sqlmock.NewRows(columns)
	if rowCount == 0 {
		return result
	}

	sqlItemValue := getSingleSqlItemValue(value)

	if tvalue.Kind() != reflect.Array && tvalue.Kind() != reflect.Slice {
		row := valueToSQLRow(sqlItemType, sqlItemValue)
		result.AddRow(row...)
		return result
	}

	for i := 0; i < rowCount; i++ {
		sqlItemValue = getSqlItemValueByIndex(value, i)
		row := valueToSQLRow(sqlItemType, sqlItemValue)
		result.AddRow(row...)
	}
	return result
}

func valueToSQLRow(itemType reflect.Type, value reflect.Value) (row []driver.Value) {
	for i := 0; i < value.NumField(); i++ {
		tag := strings.Split(itemType.Field(i).Tag.Get("json"), ",")[0]
		if tag == "" || tag == "-" {
			tag = strings.Split(itemType.Field(i).Tag.Get("db"), ",")[0]
			if tag == "" {
				continue
			}
		}

		if !IsSQLType(itemType.Field(i).Type) {
			continue
		}

		rowValue := pointers.GetUnderlyingPtrValue(value.Field(i))
		row = append(row, rowValue)

	}
	return row
}

func IsSQLType(ttype reflect.Type) bool {
	switch ttype.Kind() {
	case reflect.Int,
		reflect.String,
		reflect.Bool:
		return true
	case reflect.Ptr:
		return IsSQLType(ttype.Elem())
	case reflect.Struct:
		return ttype.Name() == reflect.TypeOf((*time.Time)(nil)).Elem().Name()
	}
	return false
}

func getSqlItemType[T any](value T) reflect.Type {
	ttype := reflect.TypeOf(value)

	switch ttype.Kind() {
	case reflect.Slice,
		reflect.Ptr,
		reflect.Interface,
		reflect.Array:
		return ttype.Elem()
	default:
		return ttype
	}
}

func getSingleSqlItemValue[T any](value T) reflect.Value {
	ttype := reflect.TypeOf(value)
	vvalue := reflect.ValueOf(value)

	switch ttype.Kind() {
	case reflect.Ptr,
		reflect.Interface:
		return vvalue.Elem()
	default:
		return vvalue
	}
}

func getSqlItemValueByIndex[T any](value T, index int) reflect.Value {
	ttype := reflect.TypeOf(value)
	vvalue := reflect.ValueOf(value)

	switch ttype.Kind() {
	case reflect.Slice,
		reflect.Array:
		return vvalue.Index(index)
	default:
		return vvalue
	}
}

func getSqlRowCount[T any](value T) int {
	ttype := reflect.TypeOf(value)
	vvalue := reflect.ValueOf(value)

	switch ttype.Kind() {
	case reflect.Slice,
		reflect.Array:
		return vvalue.Len()
	default:
		return 1
	}
}
