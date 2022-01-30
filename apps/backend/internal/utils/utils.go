package utils

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"reflect"
	"time"

	"github.com/google/uuid"
	"github.com/graph-gophers/graphql-go"
)

type ContextKey struct {
	Name string
}

func GetKey(i interface{}) string {
	return fmt.Sprintf("%T", i)
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

func GenerateHash(bytes []byte) []byte {
	hash := sha512.New512_256()
	hash.Write(bytes)
	return hash.Sum(nil)
}

func GenerateSalt() []byte {
	return GenerateHash([]byte(uuid.New().String()))
}

func GenerateSaltString() string {
	return hex.EncodeToString(GenerateSalt())
}

func GeneratePasswordHash(password string, salt string) string {
	return hex.EncodeToString(GenerateHash([]byte(password + salt)))
}

func TimeToGraphqlTime(t *time.Time) *graphql.Time {
	if t != nil {
		return &graphql.Time{Time: *t}
	}
	return nil
}
