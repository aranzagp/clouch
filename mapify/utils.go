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
