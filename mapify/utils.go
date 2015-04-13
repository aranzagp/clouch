package mapify

import (
	"reflect"
	"strings"

	"github.com/thetonymaster/clouch/utils"
)

func idTagExists(typ reflect.Type) (int, error) {
	for i := 0; i < typ.NumField(); i++ {
		strTag := typ.Field(i).Tag.Get(clouch)
		tg := utils.GetTag(strTag)
		if tg.Name == "_id" {
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

	return 0, utils.ErrNoID
}

func revTagExists(typ reflect.Type) (int, error) {
	for i := 0; i < typ.NumField(); i++ {
		strTag := typ.Field(i).Tag.Get(clouch)
		tg := utils.GetTag(strTag)
		if tg.Name == "_revs" {
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

	return 0, utils.ErrNoRevs
}

func GetID(v interface{}) (string, error) {
	typ := reflect.TypeOf(v).Elem()

	if !utils.IsStruct(typ) {
		return "", utils.ErrNotAStruct
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
