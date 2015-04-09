package mapify

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var (
	errNoID   = errors.New("No ID tag Found")
	errNoRevs = errors.New("No Rev tag Found")
)

func idTagExists(typ reflect.Type) (int, error) {
	for i := 0; i < typ.NumField(); i++ {
		strTag := typ.Field(i).Tag.Get(clouch)
		tg := getTag(strTag)
		if tg.name == "_id" {
			return i, nil
		}
	}

	for i := 0; i < typ.NumField(); i++ {
		name := typ.Field(i).Name
		name = strings.ToLower(name)

		if name == "id" {
			return i, nil
		}

	}

	return 0, errNoID
}

func revTagExists(typ reflect.Type) (int, error) {
	for i := 0; i < typ.NumField(); i++ {
		strTag := typ.Field(i).Tag.Get(clouch)
		tg := getTag(strTag)
		if tg.name == "_revs" {
			return i, nil
		}
	}

	for i := 0; i < typ.NumField(); i++ {
		name := typ.Field(i).Name
		name = strings.ToLower(name)

		if name == "revs" {
			return i, nil
		}
	}

	return 0, errNoRevs
}

func isStruct(typ reflect.Type) bool {

	if typ.Kind() != reflect.Struct {
		fmt.Println(typ.Kind())
		return false

	}
	return true
}

func GetID(v interface{}) (string, error) {
	typ := reflect.TypeOf(v).Elem()

	if !isStruct(typ) {
		return "", ErrNotAStruct
	}

	num, err := idTagExists(typ)
	if err != nil {
		return "", err
	}
	id := reflect.ValueOf(v).Elem().Field(num).String()
	return id, nil
}

func isString(i int, value reflect.Value) bool {
	return value.Field(i).Kind() == reflect.String
}

func isFloat(i int, value reflect.Value) bool {
	return value.Field(i).Kind() == reflect.Float64
}

func isSlice(i int, value reflect.Value) bool {
	return value.Field(i).Kind() == reflect.Slice
}

func isInt(i int, value reflect.Value) bool {
	return value.Field(i).Kind() == reflect.Int
}

func isPtr(i int, value reflect.Value) bool {
	return value.Field(i).Kind() == reflect.Ptr
}

func isBool(i int, value reflect.Value) bool {
	return value.Field(i).Kind() == reflect.Bool
}

func isMap(i int, value reflect.Value) bool {
	return value.Field(i).Kind() == reflect.Map
}

// func isStruct(i int, value reflect.Value) bool {
// 	return value.Field(i).Kind() == reflect.Struct
// }
