package pointers

import "reflect"

func ToPtr[T any](v T) *T {
	return &v
}

func GetUnderlyingPtrValue(vvalue reflect.Value) interface{} {
	if vvalue.Kind() == reflect.Ptr || vvalue.Kind() == reflect.Interface {
		if vvalue.IsNil() {
			return nil
		}
		return vvalue.Elem().Interface()
	}
	return vvalue.Interface()
}

func GetUnderlyingTypeRecursively(ttype reflect.Type) reflect.Type {
	if ttype.Kind() == reflect.Ptr || ttype.Kind() == reflect.Interface {
		return GetUnderlyingTypeRecursively(ttype.Elem())
	}
	return ttype
}
