package mapify

import (
	"fmt"
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
